package discord

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kolosys/discord/commands"
	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/models"
)

// commandRouter is the integrated command router.
var commandRouter *commands.Router

// Commands returns the command router for registering commands.
func (b *Bot) Commands() *commands.Router {
	b.mu.Lock()
	defer b.mu.Unlock()

	if commandRouter == nil {
		commandRouter = commands.NewRouter()
		commandRouter.SetResponder(b)
		b.setupInteractionHandler()
	}

	return commandRouter
}

// appID returns the application ID.
func (b *Bot) appID() string {
	return b.applicationID
}

// SetApplicationID sets the application ID (usually obtained from Ready event).
func (b *Bot) SetApplicationID(id string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.applicationID = id
	if commandRouter != nil {
		commandRouter.SetSyncer(b, id)
	}
}

// setupInteractionHandler registers the INTERACTION_CREATE handler.
func (b *Bot) setupInteractionHandler() {
	// Set application ID from Ready event (so syncing works immediately)
	_ = b.dispatcher.OnRaw(events.Ready, func(ctx context.Context, _ events.Type, data json.RawMessage) {
		var ready events.ReadyEvent
		if err := json.Unmarshal(data, &ready); err != nil {
			log.Printf("discord: failed to unmarshal ready event: %v", err)
			return
		}
		if ready.Application != nil && ready.Application.ID != "" {
			b.SetApplicationID(ready.Application.ID)
		}
	})

	// Handle interactions
	_ = b.dispatcher.OnRaw(events.InteractionCreate, func(ctx context.Context, _ events.Type, data json.RawMessage) {
		var interaction commands.Interaction
		if err := json.Unmarshal(data, &interaction); err != nil {
			log.Printf("discord: failed to unmarshal interaction: %v", err)
			return
		}

		log.Printf("discord: received interaction type=%d, id=%s", interaction.Type, interaction.ID)

		// Set application ID if not already set (fallback)
		if b.applicationID == "" {
			b.SetApplicationID(interaction.ApplicationID)
		}

		if err := commandRouter.HandleInteraction(ctx, &interaction); err != nil {
			log.Printf("discord: failed to handle interaction: %v", err)
		}
	})
}

// Responder interface implementation

// RespondToInteraction sends a response to an interaction.
func (b *Bot) RespondToInteraction(ctx context.Context, interactionID, token string, response *commands.InteractionResponse) error {
	err := b.rest.CreateInteractionResponse(ctx, interactionID, token, response)
	if err != nil {
		return fmt.Errorf("discord: failed to respond to interaction: %w", err)
	}
	return nil
}

// EditInteractionResponse edits the original interaction response.
func (b *Bot) EditInteractionResponse(ctx context.Context, token string, edit *commands.MessageEdit) error {
	appID := b.appID()
	if appID == "" {
		return fmt.Errorf("discord: application ID not set")
	}

	// Convert to model type
	modelEdit := &models.IncomingWebhookUpdateOptionsPartial{
		Content: edit.Content,
	}

	_, err := b.rest.EditOriginalInteractionResponse(ctx, appID, token, modelEdit)
	if err != nil {
		return fmt.Errorf("discord: failed to edit interaction response: %w", err)
	}
	return nil
}

// DeleteInteractionResponse deletes the original interaction response.
func (b *Bot) DeleteInteractionResponse(ctx context.Context, token string) error {
	appID := b.appID()
	if appID == "" {
		return fmt.Errorf("discord: application ID not set")
	}
	err := b.rest.DeleteOriginalInteractionResponse(ctx, appID, token)
	if err != nil {
		return fmt.Errorf("discord: failed to delete interaction response: %w", err)
	}
	return nil
}

// FollowupMessage sends a followup message to an interaction.
func (b *Bot) FollowupMessage(ctx context.Context, token string, message *commands.MessageCreate) (*models.Message, error) {
	appID := b.appID()
	if appID == "" {
		return nil, fmt.Errorf("discord: application ID not set")
	}
	msg, err := b.rest.CreateFollowupMessage(ctx, appID, token, message)
	if err != nil {
		return nil, fmt.Errorf("discord: failed to send followup message: %w", err)
	}
	return msg, nil
}

// CommandSyncer interface implementation

// SyncCommands syncs commands globally.
func (b *Bot) SyncCommands(ctx context.Context, appID string, cmds []commands.ApplicationCommandCreate) error {
	_, err := b.rest.BulkOverwriteGlobalCommands(ctx, appID, cmds)
	if err != nil {
		return fmt.Errorf("discord: failed to sync commands: %w", err)
	}
	return nil
}

// SyncGuildCommands syncs commands to a specific guild.
func (b *Bot) SyncGuildCommands(ctx context.Context, appID, guildID string, cmds []commands.ApplicationCommandCreate) error {
	_, err := b.rest.BulkOverwriteGuildCommands(ctx, appID, guildID, cmds)
	if err != nil {
		return fmt.Errorf("discord: failed to sync guild commands: %w", err)
	}
	return nil
}

// HTTP Interactions Endpoint

// InteractionEndpoint returns an HTTP handler for Discord's interactions endpoint.
// This can be used with the Helix server or any standard HTTP server.
func (b *Bot) InteractionEndpoint(publicKey string) http.HandlerFunc {
	pubKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		log.Printf("discord: invalid public key: %v", err)
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "invalid configuration", http.StatusInternalServerError)
		}
	}

	pubKey := ed25519.PublicKey(pubKeyBytes)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Verify signature
		signature := r.Header.Get("X-Signature-Ed25519")
		timestamp := r.Header.Get("X-Signature-Timestamp")

		if !verifyInteractionSignature(pubKey, signature, timestamp, body) {
			http.Error(w, "invalid signature", http.StatusUnauthorized)
			return
		}

		// Parse interaction
		var interaction commands.Interaction
		if err := json.Unmarshal(body, &interaction); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// Handle ping
		if interaction.Type == commands.InteractionTypePing {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"type":1}`))
			return
		}

		// Set application ID if not set
		if b.applicationID == "" {
			b.SetApplicationID(interaction.ApplicationID)
		}

		// Create response writer wrapper
		respWriter := &interactionResponseWriter{w: w}

		// Create a special responder for HTTP interactions
		httpResponder := &httpInteractionResponder{
			writer: respWriter,
			bot:    b,
		}

		// Get router and set the HTTP responder temporarily
		router := b.Commands()
		router.SetResponder(httpResponder)

		// Handle the interaction
		ctx := r.Context()
		if err := router.HandleInteraction(ctx, &interaction); err != nil {
			log.Printf("discord: failed to handle HTTP interaction: %v", err)
			if !respWriter.written {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}

		// Restore the original responder
		router.SetResponder(b)

		// If no response was written, send a default acknowledgement
		if !respWriter.written {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"type":6}`)) // Deferred update
		}
	}
}

func verifyInteractionSignature(pubKey ed25519.PublicKey, signature, timestamp string, body []byte) bool {
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	message := append([]byte(timestamp), body...)
	return ed25519.Verify(pubKey, message, sig)
}

type interactionResponseWriter struct {
	w       http.ResponseWriter
	written bool
}

func (w *interactionResponseWriter) WriteResponse(resp *commands.InteractionResponse) error {
	if w.written {
		return fmt.Errorf("response already written")
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.w.Header().Set("Content-Type", "application/json")
	_, err = w.w.Write(data)
	w.written = true
	return err
}

type httpInteractionResponder struct {
	writer *interactionResponseWriter
	bot    *Bot
	token  string
}

func (r *httpInteractionResponder) RespondToInteraction(_ context.Context, _, token string, response *commands.InteractionResponse) error {
	r.token = token
	return r.writer.WriteResponse(response)
}

func (r *httpInteractionResponder) EditInteractionResponse(ctx context.Context, token string, edit *commands.MessageEdit) error {
	return r.bot.EditInteractionResponse(ctx, token, edit)
}

func (r *httpInteractionResponder) DeleteInteractionResponse(ctx context.Context, token string) error {
	return r.bot.DeleteInteractionResponse(ctx, token)
}

func (r *httpInteractionResponder) FollowupMessage(ctx context.Context, token string, message *commands.MessageCreate) (*models.Message, error) {
	return r.bot.FollowupMessage(ctx, token, message)
}

// RegisterCommand is a convenience method to register a command.
func (b *Bot) RegisterCommand(cmd commands.Command) {
	b.Commands().Command(cmd)
}

// Slash is a convenience method to register a slash command.
func (b *Bot) Slash(name, description string, handler commands.HandlerFunc, options ...commands.Option) {
	b.Commands().Slash(name, description, handler, options...)
}

// UserContext is a convenience method to register a user context menu command.
func (b *Bot) UserContext(name string, handler commands.HandlerFunc) {
	b.Commands().UserContext(name, handler)
}

// MessageContext is a convenience method to register a message context menu command.
func (b *Bot) MessageContext(name string, handler commands.HandlerFunc) {
	b.Commands().MessageContext(name, handler)
}

// Component registers a component handler.
func (b *Bot) Component(customID string, handler commands.ComponentHandlerFunc) {
	b.Commands().Component(customID, handler)
}

// ComponentPrefix registers a component handler that matches by prefix.
func (b *Bot) ComponentPrefix(prefix string, handler commands.ComponentHandlerFunc) {
	b.Commands().ComponentPrefix(prefix, handler)
}

// SyncCommandsToDiscord syncs all registered commands with Discord.
func (b *Bot) SyncCommandsToDiscord(ctx context.Context) error {
	return b.Commands().Sync(ctx)
}

// SyncCommandsToGuild syncs all registered commands to a specific guild.
func (b *Bot) SyncCommandsToGuild(ctx context.Context, guildID string) error {
	return b.Commands().SyncGuild(ctx, guildID)
}

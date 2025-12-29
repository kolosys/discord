# main

This example demonstrates basic usage of the library.

## Source Code

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/commands"
	"github.com/kolosys/discord/components"
	"github.com/kolosys/discord/gateway"

	"github.com/kolosys/discord/examples/internal"
	_ "github.com/kolosys/discord/examples/internal"
)

func main() {
	internal.LoadEnv(".env")
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	bot, err := discord.New(&discord.Options{
		Token:   token,
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages,
		Addr:    ":8080",
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// ============================================================
	// SLASH COMMANDS
	// ============================================================

	// Simple command - no options
	bot.Slash("ping", "Check if the bot is alive", handlePing)

	// Command with required option
	bot.Slash("echo", "Echo a message back", handleEcho,
		commands.String("message", "The message to echo").Required().Build(),
	)

	// Command with multiple options
	bot.Slash("greet", "Greet a user", handleGreet,
		commands.User("user", "The user to greet").Required().Build(),
		commands.String("greeting", "Custom greeting").Build(),
	)

	// Command with choices
	bot.Slash("info", "Get information about something", handleInfo,
		commands.String("topic", "The topic to get info about").
			Required().
			Choices(
				commands.Choice{Name: "Bot", Value: "bot"},
				commands.Choice{Name: "Server", Value: "server"},
				commands.Choice{Name: "User", Value: "user"},
			).Build(),
	)

	// Command that shows buttons
	bot.Slash("confirm", "Show a confirmation dialog", handleConfirm)

	// Command that shows a select menu
	bot.Slash("choose", "Pick your favorite color", handleChoose)

	// Long-running command (demonstrates defer)
	bot.Slash("process", "Process something that takes time", handleProcess)

	// Command that opens a modal
	bot.Slash("feedback", "Submit feedback via modal", handleFeedback)

	// ============================================================
	// CONTEXT MENUS
	// ============================================================

	bot.UserContext("Get User Info", handleUserInfo)
	bot.MessageContext("Quote Message", handleQuoteMessage)

	// ============================================================
	// COMPONENT HANDLERS (buttons, selects, modals)
	// ============================================================

	// Button handlers
	bot.Component("confirm_yes", handleConfirmYes)
	bot.Component("confirm_no", handleConfirmNo)

	// Select menu handler
	bot.Component("color_select", handleColorSelect)

	// Modal submission handler
	bot.Component("feedback_modal", handleFeedbackSubmit)

	// Prefix handler - matches any custom_id starting with "action_"
	bot.ComponentPrefix("action_", handleDynamicAction)

	// ============================================================
	// MIDDLEWARE
	// ============================================================

	bot.Commands().Use(
		commands.Recover(),
		commands.Logger(log.Printf),
	)

	bot.Commands().UseComponent(
		commands.ComponentRecover(),
		commands.ComponentLogger(log.Printf),
	)

	// ============================================================
	// READY HANDLER - Sync commands
	// ============================================================

	bot.OnReady(func(ctx context.Context, e *discord.ReadyEvent) {
		log.Printf("Bot is ready as %s", e.User.Username)

		if err := bot.SyncCommandsToDiscord(ctx); err != nil {
			log.Printf("Failed to sync commands: %v", err)
		} else {
			log.Println("Commands synced successfully")
		}
	})

	// ============================================================
	// START BOT
	// ============================================================

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	log.Println("Bot is running. Press Ctrl+C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Shutting down...")
	if err := bot.Stop(ctx); err != nil {
		log.Printf("Error stopping bot: %v", err)
	}
}

// ============================================================
// COMMAND HANDLERS
// ============================================================

func handlePing(ctx *commands.Context) error {
	return ctx.Reply("🏓 Pong!")
}

func handleEcho(ctx *commands.Context) error {
	message := ctx.String("message")
	return ctx.Reply(message)
}

func handleGreet(ctx *commands.Context) error {
	user := ctx.GetUser("user")
	greeting := ctx.Options.StringDefault("greeting", "Hello")

	if user == nil {
		return ctx.ReplyEphemeral("Could not find user")
	}

	return ctx.Reply(fmt.Sprintf("%s, <@%s>! 👋", greeting, user.ID))
}

func handleInfo(ctx *commands.Context) error {
	topic := ctx.String("topic")

	switch topic {
	case "bot":
		return ctx.ReplyEmbed(&commands.Embed{
			Title:       "Bot Information",
			Description: "This is a Discord bot built with kolosys/discord",
			Color:       0x5865F2,
			Fields: []*commands.EmbedField{
				{Name: "Library", Value: "kolosys/discord", Inline: true},
				{Name: "Language", Value: "Go", Inline: true},
			},
		})
	case "server":
		if !ctx.InGuild() {
			return ctx.ReplyEphemeral("This command must be used in a server")
		}
		return ctx.Reply(fmt.Sprintf("Server ID: %s", ctx.GuildID))
	case "user":
		if ctx.User == nil {
			return ctx.ReplyEphemeral("Could not get user info")
		}
		return ctx.Reply(fmt.Sprintf("User: %s (ID: %s)", ctx.User.Username, ctx.User.ID))
	default:
		return ctx.ReplyEphemeral("Unknown topic")
	}
}

// handleConfirm - Demonstrates sending buttons with a message
func handleConfirm(ctx *commands.Context) error {
	// Using the convenience method
	return ctx.ReplyComponents(
		"Are you sure you want to proceed?",
		components.BuildComponents(
			components.SuccessButton("confirm_yes", "Yes, proceed"),
			components.DangerButton("confirm_no", "Cancel"),
		)...,
	)
}

// handleChoose - Demonstrates sending a select menu
func handleChoose(ctx *commands.Context) error {
	selectMenu := components.StringSelect("color_select").
		Placeholder("Pick your favorite color...").
		Option("Red", "red").
		Option("Green", "green").
		Option("Blue", "blue").
		OptionWithDescription("Purple", "purple", "The color of royalty").
		Build()

	return ctx.ReplyComponents(
		"What's your favorite color?",
		components.BuildComponents(selectMenu)...,
	)
}

// handleProcess - Demonstrates deferred response for long operations
func handleProcess(ctx *commands.Context) error {
	// IMPORTANT: Call Defer() within 3 seconds for long-running operations
	if err := ctx.Defer(); err != nil {
		return err
	}

	// Simulate a long-running operation
	time.Sleep(3 * time.Second)

	// Now send the actual response (you have up to 15 minutes after deferring)
	return ctx.Reply("✅ Processing complete!")
}

// handleFeedback - Demonstrates opening a modal
func handleFeedback(ctx *commands.Context) error {
	return ctx.Modal("feedback_modal", "Submit Feedback",
		components.TextInput("title", "Title").
			Placeholder("Brief summary...").
			MaxLength(100).
			Build(),
		components.TextInput("details", "Details").
			Paragraph().
			Placeholder("Tell us more...").
			MinLength(10).
			MaxLength(1000).
			Build(),
	)
}

// ============================================================
// CONTEXT MENU HANDLERS
// ============================================================

func handleUserInfo(ctx *commands.Context) error {
	if ctx.TargetUser == nil {
		return ctx.ReplyEphemeral("No user selected")
	}

	return ctx.ReplyEmbed(&commands.Embed{
		Title: ctx.TargetUser.Username,
		Fields: []*commands.EmbedField{
			{Name: "ID", Value: ctx.TargetUser.ID, Inline: true},
			{Name: "Bot", Value: fmt.Sprintf("%t", ctx.TargetUser.Bot), Inline: true},
		},
		Color: 0x57F287,
	})
}

func handleQuoteMessage(ctx *commands.Context) error {
	if ctx.TargetMessage == nil {
		return ctx.ReplyEphemeral("No message selected")
	}

	return ctx.ReplyEmbed(&commands.Embed{
		Description: fmt.Sprintf("> %s", ctx.TargetMessage.Content),
		Footer:      &commands.EmbedFooter{Text: "Quoted message"},
		Color:       0xFEE75C,
	})
}

// ============================================================
// COMPONENT HANDLERS
// ============================================================

func handleConfirmYes(ctx *commands.ComponentContext) error {
	// Update the original message (removes the buttons)
	return ctx.Update("✅ Action confirmed!")
}

func handleConfirmNo(ctx *commands.ComponentContext) error {
	return ctx.Update("❌ Action cancelled.")
}

func handleColorSelect(ctx *commands.ComponentContext) error {
	// Get the selected value
	color := ctx.SelectedValue()

	// You can also use ctx.Values for multi-select
	return ctx.Update(fmt.Sprintf("You selected: **%s** 🎨", color))
}

func handleFeedbackSubmit(ctx *commands.ComponentContext) error {
	// Get modal input values
	title := ctx.ModalValue("title")
	details := ctx.ModalValue("details")

	// Respond to the user
	return ctx.ReplyEphemeral(fmt.Sprintf(
		"Thanks for your feedback!\n\n**%s**\n%s",
		title, details,
	))
}

func handleDynamicAction(ctx *commands.ComponentContext) error {
	// CustomID will be something like "action_delete_123" or "action_edit_456"
	// You can parse it to determine what to do
	return ctx.ReplyEphemeral(fmt.Sprintf("Handled action: %s", ctx.CustomID))
}

```

## Running the Example

To run this example:

```bash
cd commands
go run main.go
```

## Expected Output

```
Hello from Proton examples!
```

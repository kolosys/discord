package commands

import (
	"context"
	"fmt"
	"log"
)

// SyncOptions configures command synchronization.
type SyncOptions struct {
	// GuildIDs limits sync to specific guilds (for development).
	// If empty, syncs globally.
	GuildIDs []string

	// DeleteUnknown removes commands not registered in the router.
	DeleteUnknown bool

	// DryRun logs what would happen without making changes.
	DryRun bool
}

// SyncResult contains the result of a sync operation.
type SyncResult struct {
	Created []string
	Updated []string
	Deleted []string
	Errors  []error
}

// SyncWithOptions syncs commands with custom options.
func (r *Router) SyncWithOptions(ctx context.Context, opts *SyncOptions) (*SyncResult, error) {
	r.mu.RLock()
	syncer := r.syncer
	appID := r.appID
	commands := r.commands
	groups := r.groups
	r.mu.RUnlock()

	if syncer == nil {
		return nil, fmt.Errorf("commands: no syncer configured")
	}

	result := &SyncResult{}

	// Build application commands
	appCmds := make([]ApplicationCommandCreate, 0, len(commands)+len(groups))

	for _, cmd := range commands {
		appCmd := commandToApplicationCommand(cmd)
		appCmds = append(appCmds, appCmd)
		result.Created = append(result.Created, cmd.Name())
	}

	for _, group := range groups {
		appCmd := groupToApplicationCommand(group)
		appCmds = append(appCmds, appCmd)
		result.Created = append(result.Created, group.Name)
	}

	if opts.DryRun {
		log.Printf("[DRY RUN] Would sync %d commands", len(appCmds))
		for _, cmd := range appCmds {
			log.Printf("[DRY RUN]   - %s (%s)", cmd.Name, commandTypeName(cmd.Type))
		}
		return result, nil
	}

	// Sync based on options
	if len(opts.GuildIDs) > 0 {
		// Guild-specific sync
		for _, guildID := range opts.GuildIDs {
			if err := syncer.SyncGuildCommands(ctx, appID, guildID, appCmds); err != nil {
				result.Errors = append(result.Errors, fmt.Errorf("guild %s: %w", guildID, err))
			} else {
				log.Printf("Synced %d commands to guild %s", len(appCmds), guildID)
			}
		}
	} else {
		// Global sync
		if err := syncer.SyncCommands(ctx, appID, appCmds); err != nil {
			result.Errors = append(result.Errors, err)
			return result, err
		}
		log.Printf("Synced %d commands globally", len(appCmds))
	}

	if len(result.Errors) > 0 {
		return result, fmt.Errorf("commands: sync completed with %d errors", len(result.Errors))
	}

	return result, nil
}

func commandTypeName(t int32) string {
	switch CommandType(t) {
	case CommandTypeChatInput:
		return "CHAT_INPUT"
	case CommandTypeUser:
		return "USER"
	case CommandTypeMessage:
		return "MESSAGE"
	default:
		return "UNKNOWN"
	}
}

// MustSync syncs commands and panics on error.
func (r *Router) MustSync(ctx context.Context) {
	if err := r.Sync(ctx); err != nil {
		panic(fmt.Sprintf("commands: failed to sync: %v", err))
	}
}

// SyncOnReady returns a handler that syncs commands when the bot is ready.
// This is useful for development to ensure commands are always up-to-date.
func (r *Router) SyncOnReady(ctx context.Context, guildIDs ...string) func() {
	return func() {
		var err error
		if len(guildIDs) > 0 {
			for _, guildID := range guildIDs {
				if err = r.SyncGuild(ctx, guildID); err != nil {
					log.Printf("commands: failed to sync to guild %s: %v", guildID, err)
				}
			}
		} else {
			if err = r.Sync(ctx); err != nil {
				log.Printf("commands: failed to sync globally: %v", err)
			}
		}
	}
}


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

	// AutoDeprecate removes commands from Discord that are not registered in code.
	// When true (default behavior for Sync/SyncGuild), unknown commands are deleted.
	// When false, existing commands not in code are preserved.
	AutoDeprecate bool

	// DryRun logs what would happen without making changes.
	DryRun bool
}

// SyncResult contains the result of a sync operation.
type SyncResult struct {
	Created    []string
	Updated    []string
	Deprecated []string
	Preserved  []string
	Errors     []error
}

// SyncWithOptions syncs commands with custom options.
func (r *Router) SyncWithOptions(ctx context.Context, opts *SyncOptions) (*SyncResult, error) {
	if opts == nil {
		opts = &SyncOptions{AutoDeprecate: true}
	}

	if opts.DryRun {
		return r.dryRunSync(ctx, opts)
	}

	combined := &SyncResult{}

	if len(opts.GuildIDs) > 0 {
		for _, guildID := range opts.GuildIDs {
			result, err := r.syncGuildInternal(ctx, guildID, opts.AutoDeprecate)
			if err != nil {
				combined.Errors = append(combined.Errors, fmt.Errorf("guild %s: %w", guildID, err))
				continue
			}
			combined.Created = append(combined.Created, result.Created...)
			combined.Updated = append(combined.Updated, result.Updated...)
			combined.Deprecated = append(combined.Deprecated, result.Deprecated...)
		}
	} else {
		result, err := r.syncInternal(ctx, opts.AutoDeprecate)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if len(combined.Errors) > 0 {
		return combined, fmt.Errorf("commands: sync completed with %d errors", len(combined.Errors))
	}

	return combined, nil
}

func (r *Router) dryRunSync(ctx context.Context, opts *SyncOptions) (*SyncResult, error) {
	r.mu.RLock()
	syncer := r.syncer
	appID := r.appID
	commands := r.commands
	groups := r.groups
	r.mu.RUnlock()

	if syncer == nil {
		return nil, fmt.Errorf("commands: no syncer configured")
	}

	// Build known command names
	knownNames := make(map[string]struct{}, len(commands)+len(groups))
	appCmds := make([]ApplicationCommandCreate, 0, len(commands)+len(groups))

	for _, cmd := range commands {
		appCmd := commandToApplicationCommand(cmd)
		appCmds = append(appCmds, appCmd)
		knownNames[cmd.Name()] = struct{}{}
	}

	for _, group := range groups {
		appCmd := groupToApplicationCommand(group)
		appCmds = append(appCmds, appCmd)
		knownNames[group.Name] = struct{}{}
	}

	mode := "auto-deprecate"
	if !opts.AutoDeprecate {
		mode = "preserve"
	}
	log.Printf("[DRY RUN] Would sync %d commands (mode: %s)", len(appCmds), mode)
	for _, cmd := range appCmds {
		log.Printf("[DRY RUN]   - %s (%s)", cmd.Name, commandTypeName(cmd.Type))
	}

	result := &SyncResult{}

	// Fetch existing commands to show what would be deprecated/preserved
	if len(opts.GuildIDs) > 0 {
		for _, guildID := range opts.GuildIDs {
			existing, err := syncer.GetGuildCommands(ctx, appID, guildID)
			if err != nil {
				log.Printf("[DRY RUN] Failed to fetch guild %s commands: %v", guildID, err)
				continue
			}
			for _, cmd := range existing {
				if _, known := knownNames[cmd.Name]; !known {
					if opts.AutoDeprecate {
						result.Deprecated = append(result.Deprecated, cmd.Name)
						log.Printf("[DRY RUN] Would deprecate %q in guild %s", cmd.Name, guildID)
					} else {
						result.Preserved = append(result.Preserved, cmd.Name)
						log.Printf("[DRY RUN] Would preserve %q in guild %s", cmd.Name, guildID)
					}
				}
			}
		}
	} else {
		existing, err := syncer.GetCommands(ctx, appID)
		if err != nil {
			log.Printf("[DRY RUN] Failed to fetch global commands: %v", err)
		} else {
			for _, cmd := range existing {
				if _, known := knownNames[cmd.Name]; !known {
					if opts.AutoDeprecate {
						result.Deprecated = append(result.Deprecated, cmd.Name)
						log.Printf("[DRY RUN] Would deprecate %q", cmd.Name)
					} else {
						result.Preserved = append(result.Preserved, cmd.Name)
						log.Printf("[DRY RUN] Would preserve %q", cmd.Name)
					}
				}
			}
		}
	}

	for name := range knownNames {
		result.Created = append(result.Created, name)
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

package discord

import "fmt"

// Module represents a self-contained feature that can register
// commands, handlers, routes, and services with a Discord bot.
//
// Modules provide a clean way to organize bot functionality into
// separate packages that can be independently developed and tested.
//
// Example usage:
//
//	type ModerationModule struct {
//	    service *moderation.Service
//	}
//
//	func (m *ModerationModule) Register(bot *discord.Bot) error {
//	    m.service = moderation.NewService()
//
//	    // Register services
//	    bot.Commands().RegisterService(m.service)
//
//	    // Register commands
//	    bot.Slash("warn", "Warn a user", m.handleWarn)
//	    bot.Slash("ban", "Ban a user", m.handleBan)
//
//	    // Register components
//	    bot.Component("ban_confirm", m.handleBanConfirm)
//	    bot.Component("ban_cancel", m.handleBanCancel)
//
//	    // Register event handlers
//	    bot.OnMessageCreate(m.handleAutomod)
//	    bot.OnMemberJoin(m.handleMemberJoin)
//
//	    // Register HTTP routes
//	    api := bot.Group("/api/moderation")
//	    api.GET("/warnings/{userID}", m.handleGetWarnings)
//	    api.POST("/warnings", m.handleCreateWarning)
//
//	    return nil
//	}
//
// See examples/modular for a complete example.
type Module interface {
	// Register registers the module's commands, handlers, routes, and services.
	// Called during bot initialization before Start().
	// Modules can access all Bot methods to register features:
	// - Commands: Slash(), UserContext(), MessageContext()
	// - Components: Component(), ComponentPrefix()
	// - Events: OnReady(), OnMessageCreate(), etc.
	// - HTTP Routes: GET(), POST(), Group(), etc.
	// - Services: Commands().RegisterService()
	// - Middleware: Commands().Use(), Commands().UseComponent(), Use()
	Register(bot *Bot) error
}

// RegisterModules registers multiple modules with the bot.
// Modules are registered in the order they are provided.
// If any module returns an error, registration stops and that error is returned.
//
// Example:
//
//	err := discord.RegisterModules(bot,
//	    &utils.Module{},
//	    &moderation.Module{},
//	    &leveling.Module{},
//	    &admin.Module{},
//	)
//	if err != nil {
//	    return fmt.Errorf("failed to register modules: %w", err)
//	}
func RegisterModules(bot *Bot, modules ...Module) error {
	for i, m := range modules {
		if err := m.Register(bot); err != nil {
			return fmt.Errorf("failed to register module %d: %w", i, err)
		}
	}
	return nil
}

// MustRegisterModules is like RegisterModules but panics on error.
// Useful for application startup where registration errors are fatal.
//
// Example:
//
//	discord.MustRegisterModules(bot,
//	    &utils.Module{},
//	    &moderation.Module{},
//	    &leveling.Module{},
//	)
func MustRegisterModules(bot *Bot, modules ...Module) {
	if err := RegisterModules(bot, modules...); err != nil {
		panic(err)
	}
}

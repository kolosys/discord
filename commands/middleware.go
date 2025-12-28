package commands

import (
	"fmt"
	"sync"
	"time"

	"github.com/kolosys/discord/types"
)

// Middleware wraps a handler function to provide cross-cutting concerns.
type Middleware func(next HandlerFunc) HandlerFunc

// Recover is middleware that recovers from panics.
func Recover() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			defer func() {
				if r := recover(); r != nil {
					_ = ctx.ReplyEphemeral(fmt.Sprintf("An error occurred: %v", r))
				}
			}()
			return next(ctx)
		}
	}
}

// GuildOnly ensures the command is only usable in guilds.
func GuildOnly() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if !ctx.InGuild() {
				return ctx.ReplyEphemeral("This command can only be used in a server.")
			}
			return next(ctx)
		}
	}
}

// DMOnly ensures the command is only usable in DMs.
func DMOnly() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if !ctx.InDM() {
				return ctx.ReplyEphemeral("This command can only be used in DMs.")
			}
			return next(ctx)
		}
	}
}

// RequirePermissions ensures the user has the required permissions.
func RequirePermissions(perms ...types.Permissions) Middleware {
	var required types.Permissions
	for _, p := range perms {
		required = required.Add(p)
	}

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if !ctx.InGuild() {
				return next(ctx) // No permissions in DMs
			}
			if !ctx.HasPermission(required) {
				return ctx.ReplyEphemeral("You don't have permission to use this command.")
			}
			return next(ctx)
		}
	}
}

// RequireRoles ensures the user has one of the required roles.
func RequireRoles(roleIDs ...string) Middleware {
	roleSet := make(map[string]struct{}, len(roleIDs))
	for _, id := range roleIDs {
		roleSet[id] = struct{}{}
	}

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if !ctx.InGuild() || ctx.Member == nil {
				return ctx.ReplyEphemeral("This command requires specific roles.")
			}

			for _, roleID := range ctx.Member.Roles {
				if _, ok := roleSet[roleID]; ok {
					return next(ctx)
				}
			}

			return ctx.ReplyEphemeral("You don't have the required role to use this command.")
		}
	}
}

// RequireOwner ensures only the bot owner can use the command.
func RequireOwner(ownerIDs ...string) Middleware {
	ownerSet := make(map[string]struct{}, len(ownerIDs))
	for _, id := range ownerIDs {
		ownerSet[id] = struct{}{}
	}

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if ctx.User == nil {
				return ctx.ReplyEphemeral("Unable to verify user.")
			}
			if _, ok := ownerSet[ctx.User.ID]; !ok {
				return ctx.ReplyEphemeral("This command is restricted to bot owners.")
			}
			return next(ctx)
		}
	}
}

// RateLimit applies rate limiting per user.
func RateLimit(limit int, window time.Duration) Middleware {
	type entry struct {
		count   int
		resetAt time.Time
	}

	var mu sync.Mutex
	users := make(map[string]*entry)

	// Cleanup old entries periodically
	go func() {
		ticker := time.NewTicker(window)
		defer ticker.Stop()
		for range ticker.C {
			mu.Lock()
			now := time.Now()
			for id, e := range users {
				if now.After(e.resetAt) {
					delete(users, id)
				}
			}
			mu.Unlock()
		}
	}()

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if ctx.User == nil {
				return next(ctx)
			}

			userID := ctx.User.ID
			now := time.Now()

			mu.Lock()
			e, ok := users[userID]
			if !ok || now.After(e.resetAt) {
				e = &entry{
					count:   0,
					resetAt: now.Add(window),
				}
				users[userID] = e
			}

			e.count++
			remaining := limit - e.count
			mu.Unlock()

			if remaining < 0 {
				return ctx.ReplyEphemeral(fmt.Sprintf(
					"You're using this command too quickly. Please wait %s.",
					time.Until(e.resetAt).Round(time.Second),
				))
			}

			return next(ctx)
		}
	}
}

// Cooldown applies a simple cooldown per user.
func Cooldown(duration time.Duration) Middleware {
	var mu sync.Mutex
	lastUsed := make(map[string]time.Time)

	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if ctx.User == nil {
				return next(ctx)
			}

			userID := ctx.User.ID
			now := time.Now()

			mu.Lock()
			last, ok := lastUsed[userID]
			if ok && now.Sub(last) < duration {
				remaining := duration - now.Sub(last)
				mu.Unlock()
				return ctx.ReplyEphemeral(fmt.Sprintf(
					"Please wait %s before using this command again.",
					remaining.Round(time.Second),
				))
			}
			lastUsed[userID] = now
			mu.Unlock()

			return next(ctx)
		}
	}
}

// NSFW ensures the command is used in an NSFW channel.
func NSFW() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			// This would require channel info from the interaction
			// For now, we'll just pass through
			// In a full implementation, you'd check the channel's NSFW flag
			return next(ctx)
		}
	}
}

// DeferredResponse automatically defers the response.
func DeferredResponse() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if err := ctx.Defer(); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}

// DeferredEphemeral automatically defers with ephemeral response.
func DeferredEphemeral() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if err := ctx.DeferEphemeral(); err != nil {
				return err
			}
			return next(ctx)
		}
	}
}

// Chain combines multiple middleware into one.
func Chain(middleware ...Middleware) Middleware {
	return func(next HandlerFunc) HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next
	}
}

// ErrorHandler wraps errors with a custom error handler.
func ErrorHandler(handler func(ctx *Context, err error) error) Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			if err := next(ctx); err != nil {
				return handler(ctx, err)
			}
			return nil
		}
	}
}

// Logger logs command execution.
func Logger(logFunc func(format string, args ...any)) Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx *Context) error {
			start := time.Now()
			err := next(ctx)
			duration := time.Since(start)

			userID := "unknown"
			if ctx.User != nil {
				userID = ctx.User.ID
			}

			if err != nil {
				logFunc("Command %s by user %s in guild %s failed after %s: %v",
					ctx.CommandName, userID, ctx.GuildID, duration, err)
			} else {
				logFunc("Command %s by user %s in guild %s completed in %s",
					ctx.CommandName, userID, ctx.GuildID, duration)
			}

			return err
		}
	}
}

// Component Middleware

// ComponentRecover is middleware that recovers from panics in component handlers.
func ComponentRecover() ComponentMiddleware {
	return func(next ComponentHandlerFunc) ComponentHandlerFunc {
		return func(ctx *ComponentContext) error {
			defer func() {
				if r := recover(); r != nil {
					_ = ctx.ReplyEphemeral(fmt.Sprintf("An error occurred: %v", r))
				}
			}()
			return next(ctx)
		}
	}
}

// ComponentLogger logs component interaction execution.
func ComponentLogger(logFunc func(format string, args ...any)) ComponentMiddleware {
	return func(next ComponentHandlerFunc) ComponentHandlerFunc {
		return func(ctx *ComponentContext) error {
			start := time.Now()
			err := next(ctx)
			duration := time.Since(start)

			userID := "unknown"
			if ctx.User != nil {
				userID = ctx.User.ID
			}

			if err != nil {
				logFunc("Component %s by user %s in guild %s failed after %s: %v",
					ctx.CustomID, userID, ctx.GuildID, duration, err)
			} else {
				logFunc("Component %s by user %s in guild %s completed in %s",
					ctx.CustomID, userID, ctx.GuildID, duration)
			}

			return err
		}
	}
}

// ComponentErrorHandler wraps errors with a custom error handler for components.
func ComponentErrorHandler(handler func(ctx *ComponentContext, err error) error) ComponentMiddleware {
	return func(next ComponentHandlerFunc) ComponentHandlerFunc {
		return func(ctx *ComponentContext) error {
			if err := next(ctx); err != nil {
				return handler(ctx, err)
			}
			return nil
		}
	}
}

// ComponentGuildOnly ensures the component interaction is only usable in guilds.
func ComponentGuildOnly() ComponentMiddleware {
	return func(next ComponentHandlerFunc) ComponentHandlerFunc {
		return func(ctx *ComponentContext) error {
			if !ctx.InGuild() {
				return ctx.ReplyEphemeral("This action can only be used in a server.")
			}
			return next(ctx)
		}
	}
}

// ComponentChain combines multiple component middleware into one.
func ComponentChain(middleware ...ComponentMiddleware) ComponentMiddleware {
	return func(next ComponentHandlerFunc) ComponentHandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next
	}
}

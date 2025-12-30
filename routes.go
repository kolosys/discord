package discord

import (
	"net/http"

	"github.com/kolosys/helix"
)

// GET registers a GET route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) GET(pattern string, handler http.HandlerFunc) {
	b.httpServer().GET(pattern, handler)
}

// POST registers a POST route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) POST(pattern string, handler http.HandlerFunc) {
	b.httpServer().POST(pattern, handler)
}

// PUT registers a PUT route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) PUT(pattern string, handler http.HandlerFunc) {
	b.httpServer().PUT(pattern, handler)
}

// PATCH registers a PATCH route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) PATCH(pattern string, handler http.HandlerFunc) {
	b.httpServer().PATCH(pattern, handler)
}

// DELETE registers a DELETE route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) DELETE(pattern string, handler http.HandlerFunc) {
	b.httpServer().DELETE(pattern, handler)
}

// OPTIONS registers an OPTIONS route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) OPTIONS(pattern string, handler http.HandlerFunc) {
	b.httpServer().OPTIONS(pattern, handler)
}

// HEAD registers a HEAD route on the HTTP server.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) HEAD(pattern string, handler http.HandlerFunc) {
	b.httpServer().HEAD(pattern, handler)
}

// Handle registers a route with a custom HTTP method.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) Handle(method, pattern string, handler http.HandlerFunc) {
	b.httpServer().Handle(method, pattern, handler)
}

// Any registers a route that matches any HTTP method.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) Any(pattern string, handler http.HandlerFunc) {
	b.httpServer().Any(pattern, handler)
}

// Static serves static files from the specified directory.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) Static(pattern, root string) {
	b.httpServer().Static(pattern, root)
}

// Group creates a route group with the specified prefix and optional middleware.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) Group(prefix string, middleware ...any) *helix.Group {
	return b.httpServer().Group(prefix, middleware...)
}

// Use adds middleware to the HTTP server that applies to all routes.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) Use(middleware ...any) {
	b.httpServer().Use(middleware...)
}

// Resource creates a resource builder for RESTful CRUD routes.
// Panics if HTTP is not enabled (Server.EnableHTTP must be true).
func (b *Bot) Resource(pattern string) *helix.ResourceBuilder {
	return b.httpServer().Resource(pattern)
}

// httpServer returns the helix.Server or panics with a helpful error.
func (b *Bot) httpServer() *helix.Server {
	s := b.Server()
	if s == nil {
		panic("discord: HTTP server not enabled - set Server.EnableHTTP = true in discord.Options")
	}
	return s
}

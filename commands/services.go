package commands

import (
	"reflect"
	"sync"
)

// ServiceRegistry holds registered services for dependency injection.
type ServiceRegistry struct {
	services map[reflect.Type]any
	mu       sync.RWMutex
}

func newServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[reflect.Type]any),
	}
}

// Register adds a service to the registry.
// The service is stored by its concrete type.
func (r *ServiceRegistry) Register(service any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t := reflect.TypeOf(service)
	r.services[t] = service
}

// Get retrieves a service by type.
func (r *ServiceRegistry) Get(t reflect.Type) (any, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.services[t]
	return s, ok
}

// Service retrieves a typed service from the context.
// Returns the zero value if the service is not registered.
//
// Example:
//
//	users := commands.Service[*UserService](ctx)
//	if users == nil {
//	    return ctx.ReplyEphemeral("Service unavailable")
//	}
func Service[T any](ctx *Context) T {
	var zero T
	if ctx.services == nil {
		return zero
	}

	t := reflect.TypeOf(zero)

	// Try exact type match first
	if s, ok := ctx.services.Get(t); ok {
		return s.(T)
	}

	// For pointer types, also try the underlying element type
	if t != nil && t.Kind() == reflect.Ptr {
		elemType := t.Elem()
		ptrType := reflect.PointerTo(elemType)
		if s, ok := ctx.services.Get(ptrType); ok {
			return s.(T)
		}
	}

	return zero
}

// ComponentService retrieves a typed service from a ComponentContext.
// Returns the zero value if the service is not registered.
func ComponentService[T any](ctx *ComponentContext) T {
	var zero T
	if ctx.services == nil {
		return zero
	}

	t := reflect.TypeOf(zero)

	if s, ok := ctx.services.Get(t); ok {
		return s.(T)
	}

	if t != nil && t.Kind() == reflect.Ptr {
		elemType := t.Elem()
		ptrType := reflect.PointerTo(elemType)
		if s, ok := ctx.services.Get(ptrType); ok {
			return s.(T)
		}
	}

	return zero
}

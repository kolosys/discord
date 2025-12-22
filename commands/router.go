package commands

import "github.com/kolosys/nova/bus"

type Command struct{}

type CommandGroup struct{}

type Middleware struct{}

type Router struct {
	commands   map[string]*Command
	groups     map[string]*CommandGroup
	middleware []Middleware
	events     *bus.EventBus
}

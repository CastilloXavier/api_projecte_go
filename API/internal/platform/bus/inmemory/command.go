package inmemory

import (
	command2 "api_project/API/kit/command"
	"context"
)

// CommandBus is an in-memory implemention of the command.Bus
type CommandBus struct {
	handler map[command2.Type]command2.Handler
}

// NewCommandBus is a in-memory implemention of the command.Bus
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handler: make(map[command2.Type]command2.Handler),
	}
}

// Dispatch implements the command.Bus interface
func (b*CommandBus) Dispatch(ctx context.Context, cmd command2.Command) error {
	handler, ok := b.handler[cmd.Type()]
	if !ok {
		return nil
	}
	return handler.Handle(ctx, cmd)
}

// Register implements the command.Bus interface
func (b *CommandBus) Register(cmdType command2.Type, handler command2.Handler) {
	b.handler[cmdType] = handler
}
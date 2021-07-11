package inmemory

import (
	"api_project/kit/command"
	"context"
)

// CommandBus is an in-memory implemention of the command.Bus
type CommandBus struct {
	handler map[command.Type]command.Handler
}

// NewCommandBus is a in-memory implemention of the command.Bus
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handler: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface
func (b* CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handler[cmd.Type()]
	if !ok {
		return nil
	}
	return handler.Handle(ctx, cmd)
}

// Register implements the command.Bus interface
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handler[cmdType] = handler
}
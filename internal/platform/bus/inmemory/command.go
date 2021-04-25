package inmemory

import (
	"context"
	"github.com/CodelyTV/go-hexagonal_http_api-course/05-02-timeouts/kit/command"
	"log"
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

	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
		}
	}()

	return nil
}

// Register implements the command.Bus interface
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handler[cmdType] = handler
}
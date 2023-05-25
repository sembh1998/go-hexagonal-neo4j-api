package inmemory

import (
	"context"
	"log"

	"github.com/sembh1998/go-hexagonal-neo4j-api/kit/command"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() command.Bus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
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

func (b *CommandBus) Register(t command.Type, h command.Handler) {
	b.handlers[t] = h
}

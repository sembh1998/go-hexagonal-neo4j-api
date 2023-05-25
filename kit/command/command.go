package command

import "context"

type Bus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Command) error
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

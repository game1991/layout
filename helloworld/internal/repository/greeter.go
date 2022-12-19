package repository

import (
	"context"
	"helloworld/dal/query"
)

// Greeter ...
type Greeter struct {
	ID    uint32
	Hello string
	User  *User
}

// GreeterInter ...
type GreeterInter interface {
	Greeter(ctx context.Context, g *Greeter) (string, error)
}

// NewGreeterInter .
func NewGreeterInter(q *query.Query) GreeterInter {
	return &greeter{q: q}
}

type greeter struct {
	q *query.Query
}

func (g *greeter) Greeter(ctx context.Context, greeter *Greeter) (string, error) {
	if greeter.User == nil {
		return "", ErrBadParam
	}
	// TODO : greeter reply

	return "hello" + greeter.User.Name, nil
}

package repository

import "context"

// Greeter ...
type Greeter struct {
	Hello string
}

// GreeterInter ...
type GreeterInter interface {
	Greeter(ctx context.Context, g *Greeter) error
}

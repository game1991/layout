package repository

import (
	"context"
	"database/sql"
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
	GetTransaction(ctx context.Context) func(fc func(tx *query.Query) error, opts ...*sql.TxOptions) error
}

// NewGreeterInter .
func NewGreeterInter(q *query.Query) GreeterInter {
	return &greeter{q: q}
}

type greeter struct {
	q *query.Query
}

// CreateGreeter .
func CreateGreeter(userinfo *User) *Greeter {
	return &Greeter{
		Hello: "您好",
		User:  userinfo,
	}
}

func (g *greeter) Greeter(ctx context.Context, greeter *Greeter) (string, error) {
	if greeter.User == nil {
		return "", ErrBadParam
	}
	// TODO : greeter reply

	return "hello" + greeter.User.Name, nil
}

func (g *greeter) GetTransaction(ctx context.Context) func(fc func(tx *query.Query) error, opts ...*sql.TxOptions) error {
	return func(fc func(tx *query.Query) error, opts ...*sql.TxOptions) error {
		// return err while paniced
		wrapFunc := func(tx *query.Query) (err error) {
			defer func() {
				if v := recover(); v != nil {
					err = v.(error)
				}
			}()

			if err := fc(tx); err != nil {
				return err
			}
			return nil
		}

		return g.q.Transaction(wrapFunc, opts...)
	}
}

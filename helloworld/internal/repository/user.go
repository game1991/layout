package repository

import (
	"context"
	"database/sql"

	"helloworld/dal/model"
	"helloworld/dal/query"
)

// Gender ...
type Gender uint8

// 性别
const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
)

// User ...
type User struct {
	ID       uint32
	Name     string
	Age      uint32
	Gender   Gender
	Birthday sql.NullTime
}

// QueryUsers ...
type QueryUsers struct {
	Name   string
	Age    uint32
	Gender Gender
}

// UserInter ...
type UserInter interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, updates map[string]interface{}) (int64, error)
	FindByID(ctx context.Context, id uint32) (*User, error)
	Search(ctx context.Context, limit, offset uint64, in *QueryUsers) ([]*User, error)
	DeleteByIDs(ctx context.Context, ids []uint32) error
}

type user struct {
	q *query.Query
}

// CreateUser ...
func CreateUser(u *User) (mu *model.User) {
	if u != nil {
		mu = &model.User{
			Name:     u.Name,
			Age:      u.Age,
			Gender:   uint32(u.Gender),
			Birthday: u.Birthday,
		}
	}
	return
}

func (u *user) Create(ctx context.Context, user *User) error {
	return u.q.User.WithContext(ctx).Create(CreateUser(user))
}

func (u *user) Update(ctx context.Context, updates map[string]interface{}) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (u *user) FindByID(ctx context.Context, id uint32) (*User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *user) Search(ctx context.Context, limit uint64, offset uint64, in *QueryUsers) ([]*User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *user) DeleteByIDs(ctx context.Context, ids []uint32) error {
	panic("not implemented") // TODO: Implement
}

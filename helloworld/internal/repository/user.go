package repository

import (
	"context"
	"database/sql"
	"time"

	"helloworld/dal/model"
	"helloworld/dal/query"

	"gorm.io/gen"
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
	ID        uint32       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string       `gorm:"column:name;not null" json:"name"`
	Age       uint32       `gorm:"column:age;not null" json:"age"`
	Gender    uint32       `gorm:"column:gender;not null" json:"gender"`
	Birthday  sql.NullTime `gorm:"column:birthday" json:"birthday"`
	CreatedAt time.Time    `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;not null" json:"updated_at"`
}

// QueryUsers ...
type QueryUsers struct {
	Name   string
	Age    uint32
	Gender Gender
}

// Condition ...
type Condition struct {
	Name     string
	Age      uint32
	Birthday sql.NullTime
}

// UserInter ...
type UserInter interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, updates map[string]interface{}) (int64, error)
	FindByID(ctx context.Context, id uint32) (*User, error)
	Search(ctx context.Context, limit, offset uint64, in *QueryUsers) ([]*User, error)
	DeleteByIDs(ctx context.Context, ids []uint32) error
	FindByCondition(ctx context.Context, cond *Condition) ([]*User, error)
}

// NewUserInter ...
func NewUserInter(q *query.Query) UserInter {
	return &user{q: q}
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

func convertToUser(u *model.User) *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Age:       u.Age,
		Gender:    u.Gender,
		Birthday:  u.Birthday,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
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

func (u *user) FindByCondition(ctx context.Context, cond *Condition) ([]*User, error) {
	// 条件查询
	if cond == nil {
		return nil, ErrBadParam
	}

	sub := u.q.User.WithContext(ctx).Scopes(
		u.withUserName(cond.Name),
		u.withAge(cond.Age),
		u.withBirthday(&cond.Birthday),
	)
	res, err := sub.Find()
	if err != nil {
		return nil, err
	}
	var ant []*User
	for _, item := range res {
		ant = append(ant, convertToUser(item))
	}
	return ant, nil
}

func (u *user) withUserName(name string) func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if name != "" {
			return tx.Where(u.q.User.Name.Eq(name))
		}
		return tx
	}
}

func (u *user) withAge(age uint32) func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if age != 0 {
			return tx.Where(u.q.User.Age.Eq(age))
		}
		return tx
	}
}

func (u *user) withBirthday(birthday *sql.NullTime) func(tx gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if birthday != nil {
			return tx.Where(u.q.User.Birthday.Eq(birthday))
		}
		return tx
	}
}

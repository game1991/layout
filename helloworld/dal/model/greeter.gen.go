// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGreeter = "greeter"

// Greeter mapped from table <greeter>
type Greeter struct {
	ID     uint32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Hello  string `gorm:"column:hello;not null" json:"hello"`
	UserID int32  `gorm:"column:user_id;not null" json:"user_id"`
}

// TableName Greeter's table name
func (*Greeter) TableName() string {
	return TableNameGreeter
}

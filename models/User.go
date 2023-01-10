package models

type User struct {
	Id   int64
	Name string

	// Created time.Time `xorm:"created"`
	// Updated time.Time `xorm:"updated"`
}

func (a *User) TableName() string {
	return "user"
}

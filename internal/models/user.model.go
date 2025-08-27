package models

import "github.com/google/uuid"

type User struct {
	id uuid.UUID `uuid:"id"`
	name string	`name`
	email string
	password string
	_type string
}

func (u *User) New(name string, email string, password string, _type string) {
	name = name
	email = email
	password = password
	_type = _type
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetType() string {
	return u._type
}

package models

import "github.com/google/uuid"

type User struct {
	id uuid.UUID `json:"id"`
	name string `json:"name"`
	email string `json:"email"`
	password string `json:"password"`
	_type string `json:"type"`
}

func (u *User) New(name string, email string, password string, _type string) {
	name = name
	email = email
	password = password
	_type = _type
}

func (u *User) GetId() uuid.UUID {
	return u.id
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

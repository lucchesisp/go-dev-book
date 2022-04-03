package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Nickname == "" {
		return errors.New("Nickname is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)
}

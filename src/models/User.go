package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/lucchesisp/go-dev-book/src/security"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(stage string) error {
	if err := u.validate(stage); err != nil {
		return err
	}

	u.format(stage)
	return nil
}

func (u *User) validate(stage string) error {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Nickname == "" {
		return errors.New("Nickname is required")
	}

	if stage == "register" && u.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); stage == "register" && err != nil {
		return errors.New("E-mail is invalid")
	}

	if stage == "register" && u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func (u *User) format(stage string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)

	if stage == "register" {
		hashedPassword, err := security.Hash(u.Password)

		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}
	
	return nil
}

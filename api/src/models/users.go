package models

import (
	"devbook-api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(operation string) error {

	if err := user.checkNullFields(operation); err != nil {

		return err
	}

	if err := user.formatFields(operation); err != nil {

		return err
	}

	return nil
}

func (user *User) checkNullFields(operation string) error {

	if user.Name == "" {

		return errors.New("name is required")
	}

	if user.Nickname == "" {

		return errors.New("nickname is required")
	}

	if user.Email == "" {

		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {

		return errors.New("invalid email format")
	}

	if user.Password == "" && operation == "creation" {

		return errors.New("password is required")
	}

	return nil
}

func (user *User) formatFields(operation string) error {

	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if operation == "creation" {

		pwHash, err := security.Hash(user.Password)

		if err != nil {

			return err
		}

		user.Password = string(pwHash)
	}

	return nil
}

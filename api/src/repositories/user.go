package repositories

import (
	"database/sql"
	"devbook-api/src/models"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {

	return &User{db: db}
}

func (userRepository User) Create(user models.User) (uint64, error) {

	statement, err := userRepository.db.Prepare(
		"INSERT INTO user (NAME, NICKNAME, EMAIL, PASSWORD) VALUES (?,?,?,?)",
	)

	if err != nil {

		return 0, err
	}

	defer statement.Close()

	response, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)

	if err != nil {

		return 0, nil
	}

	lastID, err := response.LastInsertId()

	if err != nil {

		return 0, err
	}

	return uint64(lastID), nil
}

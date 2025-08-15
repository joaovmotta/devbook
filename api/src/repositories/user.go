package repositories

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
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

func (userRepository User) FindUser(nameOrNick string) ([]models.User, error) {

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := userRepository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM user WHERE name LIKE ? OR nickname LIKE ?",
		nameOrNick,
		nameOrNick,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (userRepository User) FindUserById(ID uint64) ([]models.User, error) {

	rows, err := userRepository.db.Query(
		"SELECT id, name, nickname, email, createdAt FROM user WHERE id = ?",
		ID,
	)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {

			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

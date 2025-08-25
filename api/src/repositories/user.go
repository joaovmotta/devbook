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

		return 0, err
	}

	lastID, err := response.LastInsertId()

	if err != nil {

		return 0, err
	}

	return uint64(lastID), nil
}

func (userRepository User) Update(userId uint64, user models.User) error {

	statement, err := userRepository.db.Prepare(
		"UPDATE user SET name = ?, nickname = ?, email = ? where id = ?",
	)

	if err != nil {

		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nickname, user.Email, userId); err != nil {

		return err
	}

	return nil
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

func (userRepository User) FindUserById(ID uint64) (models.User, error) {

	var user models.User

	err := userRepository.db.QueryRow(
		"SELECT id, name, nickname, email, createdAt FROM user WHERE id = ?",
		ID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Nickname,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {

		return models.User{}, err
	}

	return user, nil
}

func (userRepository User) FindUserByEmail(email string) (models.User, error) {

	var user models.User

	err := userRepository.db.QueryRow(
		"SELECT id, password FROM user WHERE email = ?",
		email,
	).Scan(
		&user.ID,
		&user.Password,
	)

	if err != nil {

		return models.User{}, err
	}

	return user, nil
}

func (userRepository User) Delete(ID uint64) error {

	statement, err := userRepository.db.Prepare(
		"DELETE FROM user WHERE ID = ?",
	)

	if err != nil {

		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {

		return err
	}

	return nil
}

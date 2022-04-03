package repositories

import (
	"database/sql"
	"fmt"
	"github.com/lucchesisp/go-dev-book/src/models"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (models.User, error) {
	statement, err := u.db.Prepare("INSERT INTO users (" +
		"name, nickname, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return user, err
	}

	// Will execute the statement close the end of the function
	defer statement.Close()

	statementResult, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)

	user_id, err := statementResult.LastInsertId()

	if err != nil {
		return user, err
	}

	user.ID = user_id

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u users) FindByNameOrNickname(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // % nameOrNick %

	lines, err := u.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE name LIKE ? OR nickname LIKE ?", nameOrNick, nameOrNick)
	defer lines.Close()

	if err != nil {
		return nil, err
	}

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
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

func (u users) FindByID(id uint64) (models.User, error) {
	var user models.User

	lines, err := u.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE id = ?", id)
	defer lines.Close()

	if err != nil {
		return models.User{}, err
	}

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

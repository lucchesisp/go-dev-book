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

	user.ID = uint64(user_id)

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

func (u users) Update(id uint64, users models.User) (models.User, error) {

	lines, err := u.db.Query("UPDATE users SET name = ?, nickname = ? WHERE id = ?", users.Name, users.Nickname, id)
	defer lines.Close()

	if err != nil {
		return models.User{}, err
	}

	users.ID = id

	return users, nil
}

func (u users) Delete(id uint64) error {

	lines, err := u.db.Query("DELETE FROM users WHERE id = ?", id)
	defer lines.Close()

	if err != nil {
		return err
	}

	return nil
}

func (u users) FindPassword(id uint64) (string, error) {

	lines, err := u.db.Query("SELECT password FROM users WHERE id = ?", id)
	defer lines.Close()

	if err != nil {
		return "", err
	}

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.Password,
		); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (u users) UpdatePassword(id uint64, password string) error {

	lines, err := u.db.Query("UPDATE users SET password = ? WHERE id = ?", password, id)
	defer lines.Close()

	if err != nil {
		return err
	}

	return nil
}

func (u users) FindByEmail(email string) (models.User, error) {
	line, err := u.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	defer line.Close()

	if err != nil {
		return models.User{}, err
	}

	var user models.User

	if line.Next() {
		if err = line.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u users) Follow(followerID uint64, followedID uint64) error {
	statement, err := u.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(followerID, followedID); err != nil {
		return err
	}

	return nil
}

func (u users) Unfollow(followerID uint64, followedID uint64) error {
	statement, err := u.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(followerID, followedID); err != nil {
		return err
	}

	return nil
}

func (u users) FindFollowers(id uint64) ([]models.User, error) {
	lines, err := u.db.Query("SELECT id, name, nickname, email FROM users WHERE id IN (SELECT user_id FROM followers WHERE follower_id = ?)", id)
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
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u users) FindFollowing(id uint64) ([]models.User, error) {
	lines, err := u.db.Query("SELECT id, name, nickname, email FROM users WHERE id IN (SELECT follower_id FROM followers WHERE user_id = ?)", id)
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
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

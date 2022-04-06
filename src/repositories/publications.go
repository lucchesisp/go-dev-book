package repositories

import (
	"database/sql"
	"github.com/lucchesisp/go-dev-book/src/models"
)

type publications struct {
	db *sql.DB
}

func NewPublicationsRepository(db *sql.DB) *publications {
	return &publications{db}
}

func (p publications) Create(publication models.Publications, userID uint64) (models.Publications, error) {
	statement, err := p.db.Prepare("INSERT INTO publications (" +
		"content, author_id) VALUES (?, ?)")

	if err != nil {
		return publication, err
	}

	defer statement.Close()

	statementResult, err := statement.Exec(publication.Content, userID)

	publication_id, err := statementResult.LastInsertId()

	if err != nil {
		return models.Publications{}, err
	}

	publication.ID = uint64(publication_id)
	publication.AuthorID = userID

	if err != nil {
		return models.Publications{}, err
	}

	return publication, nil
}

func (p publications) FindByID(id uint64) (models.Publications, error) {
	line, err := p.db.Query("SELECT * FROM publications WHERE id = ?", id)

	defer line.Close()

	if err != nil {
		return models.Publications{}, err
	}

	publication := models.Publications{}

	if line.Next() {
		err = line.Scan(
			&publication.ID,
			&publication.Content,
			&publication.AuthorID,
			&publication.LikeCount,
			&publication.CreatedAt,
		)

		if err != nil {
			return models.Publications{}, err
		}
	}

	return publication, nil
}

func (p publications) FindAll(userID uint64) ([]models.Publications, error) {
	lines, err := p.db.Query(`SELECT DISTINCT p.*, u.nickname FROM publications p
		INNER JOIN users u ON u.id = p.author_id
    	INNER JOIN followers f on p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 desc`, userID, userID)

	defer lines.Close()

	if err != nil {
		return []models.Publications{}, err
	}

	publications := []models.Publications{}

	for lines.Next() {
		publication := models.Publications{}

		err = lines.Scan(
			&publication.ID,
			&publication.Content,
			&publication.AuthorID,
			&publication.LikeCount,
			&publication.CreatedAt,
			&publication.AuthorNickname,
		)

		if err != nil {
			return []models.Publications{}, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

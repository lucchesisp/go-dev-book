package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lucchesisp/go-dev-book/src/config"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringDatabaseConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

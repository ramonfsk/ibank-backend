package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func (db *AuthRepositoryDB) FindBy(email string, password string) (*Login, error) {
	var login Login
	err := db.client.Get(&login,
		"SELECT * FROM user WHERE email = ? AND password = ?",
		email,
		password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {
			log.Println("Error while verifying login request from database: ", err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}

	return &login, nil
}

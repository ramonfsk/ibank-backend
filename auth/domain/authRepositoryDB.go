package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/auth/errs"

	_ "github.com/go-sql-driver/mysql"
)

type AuthRepositoryDB struct {
	client *sqlx.DB
}

func (db AuthRepositoryDB) FindBy(email string, password string) (*User, *errs.AppError) {
	var user User
	err := db.client.Get(&user,
		"SELECT * FROM user WHERE email = ? AND password = ?",
		email,
		password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewInvalidCredentialsError("invalid credentials")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &user, nil
}

func NewAuthRepositoryDB(dbClient *sqlx.DB) AuthRepositoryDB {
	return AuthRepositoryDB{client: dbClient}
}

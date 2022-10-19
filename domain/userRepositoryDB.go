package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.ibm.com/rfnascimento/ibank/errs"
	"github.ibm.com/rfnascimento/ibank/logger"
)

type UserRepositoryDB struct {
	client *sqlx.DB
}

func (d UserRepositoryDB) FindAll(status string) ([]User, *errs.AppError) {
	var err error
	users := make([]User, 0)

	if status == "" {
		err = d.client.Select(&users, "SELECT * FROM user")
	} else {
		s, errParse := strconv.ParseBool(status)
		if errParse != nil {
			return nil, errs.NewInvalidParameterError("Parameter status " + status + " is invalid")
		}
		err = d.client.Select(&users, "SELECT * FROM user WHERE status = ?", s)
	}

	if err != nil {
		logger.Error("Error while querying for users table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return users, nil
}

func (d UserRepositoryDB) FindByID(id string) (*User, *errs.AppError) {
	var user User

	err := d.client.Get(&user, "SELECT * FROM user WHERE id =?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			logger.Error("Error while scanning user" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &user, nil
}

func NewUserRepositoryDB(dbClient *sqlx.DB) UserRepositoryDB {
	return UserRepositoryDB{
		client: dbClient,
	}
}

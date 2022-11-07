package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.ibm.com/rfnascimento/ibank/server/errs"
	"github.ibm.com/rfnascimento/ibank/server/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (db AccountRepositoryDB) FindAll(status string) ([]Account, *errs.AppError) {
	var err error
	accounts := make([]Account, 0)

	if status == "" {
		err = db.client.Select(&accounts, "SELECT * FROM account")
	} else {
		s, errParse := strconv.ParseBool(status)
		if errParse != nil {
			return nil, errs.NewInvalidParameterError("Parameter status " + status + " is invalid")
		}
		err = db.client.Select(&accounts, "SELECT * FROM account WHERE status = ?", s)
	}

	if err != nil {
		logger.Error("Error while querying for accounts table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return accounts, nil
}

func (db AccountRepositoryDB) FindByID(id string) (*Account, *errs.AppError) {
	var account Account

	err := db.client.Get(&account, "SELECT * FROM account WHERE id =?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		} else {
			logger.Error("Error while scanning account" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &account, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: dbClient,
	}
}

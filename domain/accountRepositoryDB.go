package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.ibm.com/rfnascimento/ibank/errs"
	"github.ibm.com/rfnascimento/ibank/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	result, err := d.client.Exec(`INSERT INTO account (user_id, opening_date, agency, number, check_digit, pin, balance, status) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		a.UserID,
		a.OpeningDate,
		a.Agency,
		a.Number,
		a.CheckDigit,
		a.PIN,
		a.Balance,
		a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while geting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.ID = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: dbClient,
	}
}

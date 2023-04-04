package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/server/errs"
	"github.com/ramonfsk/ibank-backend/server/logger"
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

	err := db.client.Get(&account, "SELECT * FROM account WHERE id = ?", id)
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

func (db AccountRepositoryDB) FindAllTransactionsByID(id string) ([]Transaction, *errs.AppError) {
	var err error
	transactions := make([]Transaction, 0)

	account, appErr := db.FindByID(id)
	if appErr != nil {
		return transactions, appErr
	}

	err = db.client.Select(&transactions,
		`SELECT * FROM transaction WHERE agency = ? AND account_number = ? AND check_digit = ?`,
		account.Agency,
		account.Number,
		account.CheckDigit)
	if err != nil {
		logger.Error("Error while querying for transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func (db AccountRepositoryDB) FindAllTransactionsByAccountIDWithPeriod(id string, startDate string, endDate string) ([]Transaction, *errs.AppError) {
	var err error
	transactions := make([]Transaction, 0)

	account, appErr := db.FindByID(id)
	if appErr != nil {
		return transactions, appErr
	}

	err = db.client.Select(&transactions,
		`SELECT * FROM transaction WHERE agency = ? AND account_number = ? AND check_digit = ? AND created_at BETWEEN ? AND ?`,
		account.Agency,
		account.Number,
		account.CheckDigit,
		startDate+" 00:00:00",
		endDate+" 23:59:00")
	if err != nil {
		logger.Error("Error while querying for transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: dbClient,
	}
}

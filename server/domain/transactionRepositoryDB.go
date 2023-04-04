package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/server/errs"
	"github.com/ramonfsk/ibank-backend/server/logger"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (db TransactionRepositoryDB) FindAll() ([]Transaction, *errs.AppError) {
	var err error
	transactions := make([]Transaction, 0)

	err = db.client.Select(&transactions,
		`SELECT * FROM transaction`)
	if err != nil {
		logger.Error("Error while querying for transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func (db TransactionRepositoryDB) FindByID(id string) (*Transaction, *errs.AppError) {
	var transaction Transaction

	err := db.client.Get(&transaction, "SELECT * FROM transaction WHERE id =?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Transaction not found")
		} else {
			logger.Error("Error while scanning transaction" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &transaction, nil
}

func (db TransactionRepositoryDB) FindAllByAccountID(accountId string) ([]Transaction, *errs.AppError) {
	var err error
	transactions := make([]Transaction, 0)

	err = db.client.Select(&transactions,
		`SELECT * FROM transaction WHERE account_id = ?`, accountId)
	if err != nil {
		logger.Error("Error while querying for transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func (db TransactionRepositoryDB) FindAllByAccountIDWithPeriod(accountId string, startDate string, endDate string) ([]Transaction, *errs.AppError) {
	var err error
	transactions := make([]Transaction, 0)

	err = db.client.Select(&transactions,
		`SELECT * FROM transaction WHERE account_id = ? AND created_at BETWEEN ? AND ?`,
		accountId,
		startDate,
		endDate)
	if err != nil {
		logger.Error("Error while querying for transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func (db TransactionRepositoryDB) RegisterNewTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// staring the database transaction
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// insert bank transaction
	result, _ := tx.Exec(`INSERT INTO transaction (bank_id, agency, account_number, check_digit, type, value)
		VALUES (?, ?, ?, ?, ?, ?)`,
		t.BankID,
		t.Agency,
		t.Number,
		t.CheckDigit,
		t.Type,
		t.Value)
	// update account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE account SET balance = balance - ? WHERE id = ?`, t.Value, t.AccountID)
	} else {
		_, err = tx.Exec(`UPDATE account SET balance = balance + ? WHERE id = ?`, t.Value, t.AccountID)
	}
	// in case of error rollback, and changes from both the tables will be revert
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction ID: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	t.ID = strconv.FormatInt(transactionID, 10)

	return &t, nil
}

func (db TransactionRepositoryDB) FindAccount(accountId string) (*Account, *errs.AppError) {
	var account Account

	err := db.client.Get(&account,
		`SELECT * FROM account WHERE id = ?`, accountId)
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

func (db TransactionRepositoryDB) FindAccountWithoutID(agency string, number string, checkDigit string) (*Account, *errs.AppError) {
	var account Account

	err := db.client.Get(&account,
		`SELECT * FROM account WHERE agency = ? AND number = ? AND check_digit = ?`,
		agency,
		number,
		checkDigit)
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

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client: dbClient}
}

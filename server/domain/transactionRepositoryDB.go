package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.ibm.com/rfnascimento/ibank/server/errs"
	"github.ibm.com/rfnascimento/ibank/server/logger"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

// func (d TransactionRepositoryDB) GetByPeriod(days int) ([]Transaction, *errs.AppError) {}

func (d TransactionRepositoryDB) GetAllTransactionsByAccount(t Transaction) ([]Transaction, *errs.AppError) {
	var err error
	transactions := make([]Transaction, 0)

	err = d.client.Select(&transactions,
		`SELECT * FROM transaction WHERE agency = ? AND number = ? AND check_digit = ?`,
		t.Agency,
		t.Number,
		t.CheckDigit)

	if err != nil {
		logger.Error("Error while querying for transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func (d TransactionRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// staring the database transaction
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// insert bank transaction
	result, _ := tx.Exec(`INSERT INTO transaction (bank_id, agency, number, check_digit, type, value)
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
	// getting account information from the accounts table
	account, appErr := d.FindAccount(t.Agency, t.Number, t.CheckDigit)
	if appErr != nil {
		return nil, appErr
	}

	t.ID = strconv.FormatInt(transactionID, 10)
	// updating the transaction struct with the last balance
	t.Value = account.Balance
	return &t, nil
}

func (d TransactionRepositoryDB) FindAccount(agency string, number string, checkDigit string) (*Account, *errs.AppError) {
	var account Account

	err := d.client.Get(&account,
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
	return TransactionRepositoryDB{
		client: dbClient,
	}
}

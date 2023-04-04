package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
	"github.com/ramonfsk/ibank-backend/server/logger"
	"github.com/ramonfsk/ibank-backend/server/utils"
)

type UserRepositoryDB struct {
	client *sqlx.DB
}

func (db UserRepositoryDB) FindAll(status string) ([]User, *errs.AppError) {
	var err error
	users := make([]User, 0)

	if status == "" {
		err = db.client.Select(&users, "SELECT * FROM user")
	} else {
		s, errParse := strconv.ParseBool(status)
		if errParse != nil {
			return nil, errs.NewInvalidParameterError("Parameter status " + status + " is invalid")
		}
		err = db.client.Select(&users, "SELECT * FROM user WHERE status = ?", s)
	}

	if err != nil {
		logger.Error("Error while querying for users table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return users, nil
}

func (db UserRepositoryDB) FindByID(id string) (*User, *errs.AppError) {
	var user User

	err := db.client.Get(&user, "SELECT * FROM user WHERE id =?", id)
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

func (db UserRepositoryDB) RegisterNewUser(user dto.UserRequest, generatedAccount Account) (*Account, *errs.AppError) {
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	birthdate, err := utils.StringToDate(user.Birthdate)
	if err != nil {
		logger.Error("Error to parse date" + err.Error())
		return nil, errs.NewParseError("Error to parse date" + err.Error())
	}

	resultInsertUser, err := tx.Exec(`INSERT INTO user 
	(name, birthdate, password, email, document, phone, is_admin)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		user.Name,
		birthdate,
		user.Password,
		user.Email,
		user.Document,
		user.Phone,
		user.IsAdmin)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	idUser, err := resultInsertUser.LastInsertId()
	if err != nil {
		logger.Error("Error while geting last insert id for new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	resultInsertAccount, err := tx.Exec(`INSERT INTO account 
	(user_id, agency, number, check_digit)
	VALUES (?, ?, ?, ?)`,
		strconv.FormatInt(idUser, 10),
		generatedAccount.Agency,
		generatedAccount.Number,
		generatedAccount.CheckDigit)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	accountId, err := resultInsertAccount.LastInsertId()
	if err != nil {
		logger.Error("Error while geting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &Account{
		ID:         strconv.FormatInt(accountId, 10),
		Agency:     generatedAccount.Agency,
		Number:     generatedAccount.Number,
		CheckDigit: generatedAccount.CheckDigit,
	}, nil
}

func NewUserRepositoryDB(dbClient *sqlx.DB) UserRepositoryDB {
	return UserRepositoryDB{client: dbClient}
}

package domain

import (
	"database/sql"
	"time"

	"github.com/go-kit/kit/auth/jwt"
	"github.ibm.com/rfnascimento/ibank/auth/errs"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	ID      sql.NullInt64 `db:"id"`
	Email   string        `db:"email"`
	isAdmin bool          `db:"is_admin"`
}

func (l Login) GenerateToken() (*string, errs.AppError) {
	var claims jwt.ClaimsFactory
	if l.isAdmin {
		claims = l.claimsForAdmin()
	} else {
		claims = l.claimsForUser()
	}

	return token, nil
}

func (l Login) claimsForAdmin() jwt.ClaimsFactory {
	return &jwt.ClaimsFactory{
		"email": l.Email,
		"exp":   time.Now().Add(TOKEN_DURATION).Unix(),
	}
}

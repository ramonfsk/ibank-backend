package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ramonfsk/ibank-backend/auth/errs"
)

const TOKEN_DURATION = time.Hour

type JWTData struct {
	jwt.RegisteredClaims
	CustomClaims map[string]string `json:"custom_claims"`
}

func (u User) GenerateToken() (*string, *errs.AppError) {
	claims := JWTData{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(TOKEN_DURATION))),
		},
		CustomClaims: map[string]string{
			"id":       u.ID,
			"email":    u.Email,
			"password": u.Password,
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var hmacSecrecKey []byte

	if keyData, err := os.ReadFile("secret/hmacSecretKey"); err == nil {
		hmacSecrecKey = keyData
	} else {
		return nil, errs.NewReadingEnvironmentFileError(err.Error())
	}

	if signedTokenAsString, err := tokenString.SignedString(hmacSecrecKey); err != nil {
		return nil, errs.NewInvalidCredentialsError(err.Error())
	} else {
		return &signedTokenAsString, nil
	}
}

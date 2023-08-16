package app

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramonfsk/ibank-backend/auth/dto"
	"github.com/ramonfsk/ibank-backend/auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&loginRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		token, err := ah.service.Login(loginRequest)
		if err != nil {
			c.AbortWithStatusJSON(err.Code, err.AsMessage())
		} else {
			c.JSON(http.StatusOK, token)
		}
	}
}

func (ah *AuthHandler) Verify(c *gin.Context) {
	var authRequest dto.AuthRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&authRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(200, gin.H{"isAuthorized": true})
	}
}

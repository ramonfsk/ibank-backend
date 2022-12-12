package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/auth/dto"
	"github.ibm.com/rfnascimento/ibank/auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (ah *AuthHandler) NotImplementedHanlder(c *gin.Context) {
	fmt.Fprint(c, "Hanlder not implemented...")
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&loginRequest); err != nil {

	} else {
		token, err := ah.service.Login(loginRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.AsMessage())
		}
	}
}

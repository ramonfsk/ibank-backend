package app

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/server/errs"
	"github.ibm.com/rfnascimento/ibank/server/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) getAllAccounts(c *gin.Context) {
	status := c.Query("status")
	users, err := ah.service.GetAllAccounts(status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err.AsMessage())
	} else {
		c.JSON(http.StatusOK, users)
	}
}
func (ah *AccountHandler) getAccount(c *gin.Context) {
	regex, _ := regexp.Compile(`[0-9]+`)

	id := c.Param("id_account")
	if !regex.MatchString(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	} else {
		user, err := ah.service.GetAccount(id)
		if err != nil {
			c.JSON(http.StatusBadGateway, err.AsMessage())
		}

		c.JSON(http.StatusOK, user)
	}
}

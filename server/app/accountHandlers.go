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
	// TODO:
}
func (ah *AccountHandler) getAccount(c *gin.Context) {
	regex, _ := regexp.Compile(`[0-9]+`)

	userID := c.Param(`id_account`)
	if !regex.MatchString(userID) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: `invalid id`}).AsMessage())
	} else {
		// TODO:
	}
}

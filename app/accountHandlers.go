package app

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/dto"
	"github.ibm.com/rfnascimento/ibank/errs"
	"github.ibm.com/rfnascimento/ibank/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) newAccount(c *gin.Context) {
	regex, _ := regexp.Compile(`[0-9]+`)

	userID := c.Param(`id`)
	if !regex.MatchString(userID) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: `invalid id`}).AsMessage())
	} else {
		var request dto.NewAccountRequest
		err := json.NewDecoder(c.Request.Body).Decode(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			request.UserID = userID
			account, appError := ah.service.NewAccount(request)
			if appError != nil {
				c.JSON(appError.Code, appError.Message)
			} else {
				c.JSON(http.StatusCreated, account)
			}
		}
	}
}

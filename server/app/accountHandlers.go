package app

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ramonfsk/ibank/server/errs"
	"github.com/ramonfsk/ibank/server/service"
	"github.com/ramonfsk/ibank/server/utils"
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
	regex := regexp.MustCompile(utils.DIGITSONLYREGEX)

	id := c.Param("id")
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

func (ah *AccountHandler) getAllTransactionAccount(c *gin.Context) {
	regex := regexp.MustCompile(utils.DIGITSONLYREGEX)

	id := c.Param("id")
	if !regex.MatchString(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	} else {
		transactions, err := ah.service.GetAllTransactionByAccount(id)
		if err != nil {
			c.JSON(http.StatusBadGateway, err.AsMessage())
		}

		c.JSON(http.StatusOK, transactions)
	}
}

func (ah *AccountHandler) getAllTransactionAccountWithPeriod(c *gin.Context) {
	regexId := regexp.MustCompile(utils.DIGITSONLYREGEX)
	regexDate := regexp.MustCompile(utils.DATEFORMATREGEX)

	id := c.Param("id")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if !regexId.MatchString(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	} else if !regexDate.MatchString(startDate) || !regexDate.MatchString(endDate) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid dates"}).AsMessage())
	} else {
		transactions, err := ah.service.GetAllTransactionsByAccountWithPeriod(id, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusBadGateway, err.AsMessage())
		}

		c.JSON(http.StatusOK, transactions)
	}
}

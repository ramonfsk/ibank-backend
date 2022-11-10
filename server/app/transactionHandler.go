package app

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/server/dto"
	"github.ibm.com/rfnascimento/ibank/server/errs"
	"github.ibm.com/rfnascimento/ibank/server/service"
	"github.ibm.com/rfnascimento/ibank/server/utils"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (th *TransactionHandler) getAllTransactions(c *gin.Context) {
	transactions, err := th.service.GetAllTransactions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err.AsMessage())
	} else {
		c.JSON(http.StatusOK, transactions)
	}
}

func (th *TransactionHandler) getTransaction(c *gin.Context) {
	regex, _ := regexp.Compile(utils.DIGITSONLYREGEX)

	id := c.Param("id")
	if !regex.MatchString(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	} else {
		user, err := th.service.GetTransaction(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, err.AsMessage())
		}

		c.JSON(http.StatusOK, user)
	}
}

func (th *TransactionHandler) newTransaction(c *gin.Context) {
	var request dto.NewTransactionRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	} else {
		transaction, appErr := th.service.NewTransaction(request)
		if appErr != nil {
			c.AbortWithStatusJSON(appErr.Code, appErr.AsMessage())
		} else {
			c.JSON(http.StatusOK, transaction)
		}
	}
}

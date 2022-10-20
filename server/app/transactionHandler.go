package app

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/server/dto"
	"github.ibm.com/rfnascimento/ibank/server/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

// func (th *TransactionHandler) getByPeriod(c *gin.Context) {}

func (th *TransactionHandler) makeTransaction(c *gin.Context) {
	var request dto.TransactionRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	} else {
		action := c.Param("action")
		if action == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		} else if action == "/create" {
			transaction, appErr := th.service.MakeTransaction(request)
			if appErr != nil {
				c.AbortWithStatusJSON(appErr.Code, appErr.AsMessage())
			} else {
				c.JSON(http.StatusOK, transaction)
			}
		} else {
			transactions, appErr := th.service.GetAllTransactionsByAccount(request)
			if appErr != nil {
				c.AbortWithStatusJSON(appErr.Code, appErr.AsMessage())
			} else {
				c.JSON(http.StatusOK, transactions)
			}
		}
	}
}

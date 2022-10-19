package app

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/dto"
	"github.ibm.com/rfnascimento/ibank/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

// func (th *TransactionHandler) getByPeriod(c *gin.Context) {
// 	transaction, err := th.service.GetByPeriod("10")
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, err.AsMessage())
// 		return
// 	}

// 	c.JSON(http.StatusOK, transaction)
// }

func (th *TransactionHandler) makeTransaction(c *gin.Context) {
	// get the account_id and destination_account_id from body
	var request dto.TransactionRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	} else {
		// make transaction
		action := c.Param("create")
		if action == "create" {
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

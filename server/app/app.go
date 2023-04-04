package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/server/domain"
	"github.com/ramonfsk/ibank-backend/server/service"
)

func Start() {
	router := gin.Default()
	// setup DB connection
	dbClient := getDBClient()
	// wiring
	uh := UserHandler{service: service.NewUserService(domain.NewUserRepositoryDB(dbClient))}
	ah := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDB(dbClient))}
	th := TransactionHandler{service: service.NewTransactionService(domain.NewTransactionRepositoryDB(dbClient))}
	// define routes
	router.GET("/users", uh.getAllUsers)
	router.GET("/users/:id", uh.getUser)
	router.POST("/users", uh.newUser)

	router.GET("/accounts", ah.getAllAccounts)
	router.GET("/accounts/:id", ah.getAccount)
	router.GET("/accounts/:id/transactions", ah.getAllTransactionAccount)
	router.GET("/accounts/:id/transactionswithperiod", ah.getAllTransactionAccountWithPeriod)

	router.GET("/transactions", th.getAllTransactions)
	router.GET("/transactions/:id", th.getTransaction)
	router.POST("/transactions", th.newTransaction)
	// starting server
	router.Run(":8010")
}

func getDBClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ibank")
	if err != nil {
		panic(err)
	}
	// See Important settings section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

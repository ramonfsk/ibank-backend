package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.ibm.com/rfnascimento/ibank/domain"
	"github.ibm.com/rfnascimento/ibank/service"
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
	router.POST("/users/:id/account", ah.newAccount)
	router.POST("/transactions/*action", th.makeTransaction)
	router.POST("/transactions", th.makeTransaction)
	// starting server
	router.Run(":8000")
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

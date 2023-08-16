package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/server/domain"
	"github.com/ramonfsk/ibank-backend/server/logger"
	"github.com/ramonfsk/ibank-backend/server/service"
	"golang.org/x/sync/errgroup"
)

var (
	groupError errgroup.Group
)

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

func configureHandlers() (authHandler AuthMiddlewareHandler, uh UserHandler, ah AccountHandler, th TransactionHandler) {
	// setup DB connection
	dbClient := getDBClient()
	// wiring
	return AuthMiddlewareHandler{repository: domain.RemoteAuthRepository{}},
		UserHandler{service: service.NewUserService(domain.NewUserRepositoryDB(dbClient))},
		AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDB(dbClient))},
		TransactionHandler{service: service.NewTransactionService(domain.NewTransactionRepositoryDB(dbClient))}
}

func buildEngine() http.Handler {
	engine := gin.Default()
	// wiring
	authHandler, uh, ah, th := configureHandlers()
	// version group
	versionGroup := engine.Group("/v1")
	{
		versionGroup.POST("/users", uh.newUser)
		// auth group middleware
		authGroupMiddleware := engine.Group(versionGroup.BasePath()).Use(authHandler.authorizationMiddlewareHandler)
		{
			// define used routes
			authGroupMiddleware.GET("/users/:id", uh.getUser)
			authGroupMiddleware.GET("/accounts/:id", ah.getAccount)
			authGroupMiddleware.GET("/accounts/:id/transactions", ah.getAllTransactionAccount)
			authGroupMiddleware.GET("/accounts/:id/transactionswithperiod", ah.getAllTransactionAccountWithPeriod)
			authGroupMiddleware.GET("/transactions/:id", th.getTransaction)
			authGroupMiddleware.POST("/transactions", th.newTransaction)
		}
	}

	return engine
}

func buildAdminEngine() http.Handler {
	engine := gin.Default()
	// wiring
	authHandler, uh, ah, th := configureHandlers()
	// version group
	versionGroup := engine.Group("/v1")
	{
		// auth group middleware
		authGroupMiddleware := engine.Group(versionGroup.BasePath()).Use(authHandler.authorizationMiddlewareHandler)
		{
			// define admin routes
			authGroupMiddleware.GET("/users", uh.getAllUsers)
			authGroupMiddleware.GET("/accounts", ah.getAllAccounts)
			authGroupMiddleware.GET("/transactions", th.getAllTransactions)
		}
	}

	return engine
}

func Start() {
	server := &http.Server{
		Addr:         ":8010",
		Handler:      buildEngine(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	adminServer := &http.Server{
		Addr:         ":8012",
		Handler:      buildAdminEngine(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	groupError.Go(func() error {
		return server.ListenAndServe()
	})

	groupError.Go(func() error {
		return adminServer.ListenAndServe()
	})

	if err := groupError.Wait(); err != nil {
		logger.Error(err.Error())
	}
}

package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ramonfsk/ibank-backend/auth/domain"
	"github.com/ramonfsk/ibank-backend/auth/service"
)

func Start() {
	sanityCheck()

	router := gin.Default()

	dbClient := getDBClient()

	ah := AuthHandler{service: service.NewAuthService(domain.NewAuthRepositoryDB(dbClient))}

	router.POST("/auth/login", ah.Login)
	router.POST("/auth/register", ah.Register)
	router.GET("/auth/verify", ah.Verify)

	router.Run(":8011")
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

func sanityCheck() {

}

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

	engine := gin.Default()

	dbClient := getDBClient()

	ah := AuthHandler{service: service.NewAuthService(domain.NewAuthRepositoryDB(dbClient))}

	versionGroup := engine.Group("/v1")
	{
		versionGroup.POST("/auth/login", ah.Login)
		versionGroup.POST("/auth/verify", ah.Verify)
	}

	engine.Run(":8011")
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

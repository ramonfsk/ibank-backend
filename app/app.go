package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// define routes
	router.GET("/greet", greet)
	router.GET("/customers", getAllCostumers)
	router.GET("/customers/:customer_id", getCustomer)
	// starting server
	router.Run(":8000")
}

package app

import (
	"github.com/gin-gonic/gin"
)

type Costumer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(c *gin.Context) {
	c.String(200, "Hello world!")
}

func getAllCostumers(c *gin.Context) {
	costumers := []Costumer{
		{Name: "Ramon", City: "NY", Zipcode: "123"},
		{Name: "Jonh", City: "Atlanta", Zipcode: "234"},
	}

	if c.GetHeader("Content-Type") == "application/xml" {
		c.Request.Header.Add("Content-Type", "application/xml")
		c.XML(200, costumers)
	} else {
		c.Request.Header.Add("Content-Type", "application/json")
		c.JSON(200, costumers)
	}
}

func getCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	c.JSON(200, customerID)
}

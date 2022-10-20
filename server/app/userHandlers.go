package app

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.ibm.com/rfnascimento/ibank/server/errs"
	"github.ibm.com/rfnascimento/ibank/server/service"
)

type UserHandler struct {
	service service.UserService
}

func (uh *UserHandler) getAllUsers(c *gin.Context) {
	status := c.Query("status")
	users, err := uh.service.GetAllUsers(status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err.AsMessage())
	} else {
		if c.GetHeader("Content-Type") == "application/xml" {
			c.Request.Header.Add("Content-Type", "application/xml")
			c.XML(http.StatusOK, users)
		} else {
			c.Request.Header.Add("Content-Type", "application/json")
			c.JSON(http.StatusOK, users)
		}
	}
}

func (uh *UserHandler) getUser(c *gin.Context) {
	regex, _ := regexp.Compile(`[0-9]+`)

	id := c.Param("id")
	if !regex.MatchString(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	} else {
		user, err := uh.service.GetUser(id)
		if err != nil {
			c.JSON(http.StatusBadGateway, err.AsMessage())
		}

		c.JSON(http.StatusOK, user)
	}
}

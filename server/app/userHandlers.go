package app

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
	"github.com/ramonfsk/ibank-backend/server/service"
	"github.com/ramonfsk/ibank-backend/server/utils"
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
		c.JSON(http.StatusOK, users)
	}
}

func (uh *UserHandler) getUser(c *gin.Context) {
	regex := regexp.MustCompile(utils.DIGITSONLYREGEX)

	id := c.Param("id")
	if !regex.MatchString(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	} else {
		user, err := uh.service.GetUser(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, err.AsMessage())
		}

		c.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) newUser(c *gin.Context) {
	var request dto.UserRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	} else {
		user, appErr := uh.service.NewUser(request)
		if appErr != nil {
			c.AbortWithStatusJSON(appErr.Code, appErr.AsMessage())
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

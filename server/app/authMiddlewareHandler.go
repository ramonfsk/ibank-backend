package app

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ramonfsk/ibank-backend/server/domain"
	"github.com/ramonfsk/ibank-backend/server/errs"
	"github.com/ramonfsk/ibank-backend/server/utils"
)

type AuthMiddlewareHandler struct {
	repository domain.AuthRepository
}

func (ah AuthMiddlewareHandler) authorizationMiddlewareHandler(c *gin.Context) {
	regex := regexp.MustCompile(utils.DIGITSONLYREGEX)
	id := c.Param("id")
	if !regex.MatchString(id) {
		c.Abort()
		c.AbortWithStatusJSON(http.StatusBadRequest, (&errs.AppError{Message: "invalid id"}).AsMessage())
	}

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.Abort()
		c.AbortWithStatusJSON(http.StatusUnauthorized, (&errs.AppError{Message: "Missing token"}).AsMessage())
	} else {
		isAuthorized, err := ah.repository.IsAuthorized(token, id, c.FullPath(), c.Request.Method)
		if err != nil {
			c.Abort()
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		} else if !isAuthorized {
			c.Abort()
			c.AbortWithStatusJSON(http.StatusForbidden, (&errs.AppError{Message: "Unauthorized"}).AsMessage())
		} else {
			c.Next()
		}
	}
}

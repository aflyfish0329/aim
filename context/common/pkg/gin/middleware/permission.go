package middleware

import (
	"github.com/gin-gonic/gin"
)

func NewPermissionMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: get user from gin.Context

		// TODO: validate user permission through access api

		c.Next()
	}
}

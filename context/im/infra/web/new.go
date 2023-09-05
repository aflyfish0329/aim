package web

import (
	"github.com/gin-gonic/gin"
)

func NewWeb() *gin.Engine {
	c := gin.New()

	return c
}

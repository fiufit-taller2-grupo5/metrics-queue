package api

import "github.com/gin-gonic/gin"

type Controller interface {
	SetUp(router gin.IRouter)
}

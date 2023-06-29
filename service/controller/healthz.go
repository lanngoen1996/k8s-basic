package controller

import "github.com/gin-gonic/gin"

type HealthzI interface {
	Ping(c *gin.Context)
}

type HealthzController struct{}

func NewHealthzController() HealthzI {
	return &HealthzController{}
}

func (hzc *HealthzController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

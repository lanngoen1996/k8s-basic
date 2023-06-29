package route

import (
	"go-service/controller"

	"github.com/gin-gonic/gin"
)

func HealthzRouter(r *gin.Engine) {
	c := controller.NewHealthzController()
	r.GET("/ping", c.Ping)
}

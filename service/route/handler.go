package route

import "github.com/gin-gonic/gin"

type HandlerRoute struct{}

func NewHandler(r *gin.Engine) {
	HealthzRouter(r)
}

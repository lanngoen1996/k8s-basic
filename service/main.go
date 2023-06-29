package main

import (
	"go-service/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	route.NewHandler(r)

	r.Run(":3333")
}

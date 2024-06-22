package main

import (
	"JVT_Tech/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	h := gin.New()
	routes.Rout(h)
	h.Run(":8080")
}

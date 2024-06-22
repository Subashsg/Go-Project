package routes

import (
	"JVT_Tech/controller"

	"github.com/gin-gonic/gin"
)

func Rout(r *gin.Engine) {
	r.POST("/developer", controller.Create_developer())
}

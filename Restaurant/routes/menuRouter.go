package routes

import (
	"Restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(f *gin.Engine) {

	f.POST("/menus", controllers.CreateMenu())
	f.PATCH("/menus/:menu_id", controllers.UpdateFood())
	f.GET("/menus/:menu_id", controllers.GetMenu())
}

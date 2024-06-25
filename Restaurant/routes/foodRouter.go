package routes

import (
	"Restaurant/controllers"

	"github.com/gin-gonic/gin"
)

var FoodRoutes = func(r *gin.Engine) {

	r.POST("/foods", controllers.CreateFood())
	r.PATCH("/foods/:food_id", controllers.UpdateFood())
	r.GET("/foods/:food_id", controllers.GetFood())
}

package main

import (
	"Restaurant/routes"

	"github.com/gin-gonic/gin"
)

//var foodCollection *mongo.Collection = database.DatabaseTableCreation(database.Global, "food")

func main() {
	router := gin.New()
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)

	router.Run()

}

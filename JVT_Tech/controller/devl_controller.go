package controller

import (
	"JVT_Tech/database"
	"JVT_Tech/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var ret *mongo.Collection = func() *mongo.Collection {
	r := database.Global.Database("coder_range").Collection("developer")
	return r

}()

func Create_developer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var developer model.Developer
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
		if err := c.BindJSON(&developer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ret.InsertOne(ctx, developer)
		rt, _ := ret.InsertOne(ctx, developer)
		c.JSON(200, rt)

	}
}

package controllers

import (
	"Restaurant/database"
	"Restaurant/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ret *mongo.Collection = database.DatabaseTableCreation(database.Global, "food")

/*var ret *mongo.Collection = func() *mongo.Collection {
	r := database.Global.Database("restaurant").Collection("food")
	return r
}()*/

func CreateFood() gin.HandlerFunc {

	return func(c *gin.Context) {

		var food models.Food

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		// Input of API
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Responce of API
			return
		}
		var validate *validator.Validate // Validate of API

		validate = validator.New() //New return communication

		validateErr := validate.Struct(food)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Id = primitive.NewObjectID()
		food.Food_id = food.Id.Hex()

		rt, _ := ret.InsertOne(ctx, food)
		c.JSON(http.StatusOK, rt)
	}
}

func UpdateFood() gin.HandlerFunc {

	return func(c *gin.Context) {

		var food models.Food

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		foodId := c.Param("food_id")

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D

		updateObj = append(updateObj, bson.E{"name", food.Name})

		//updateObj = append(updateObj, bson.E{"name", food.Food_image})

		//updateObj = append(updateObj, bson.E{"name", food.Price})

		upsert := true

		ops := options.UpdateOptions{
			Upsert: &upsert,
		}

		filter := bson.M{"food_id": foodId}

		result, err := ret.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &ops)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, result)
	}
}

func GetFood() gin.HandlerFunc {

	return func(c *gin.Context) {

		var food models.Food

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		foodId := c.Param("food_id")

		err := ret.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

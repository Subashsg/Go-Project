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

var ret1 *mongo.Collection = database.DatabaseTableCreation(database.Global, "menu")

func CreateMenu() gin.HandlerFunc {

	return func(c *gin.Context) {

		var menu models.Menu

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var validate *validator.Validate

		validate = validator.New()

		validateErr := validate.Struct(menu)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		ret1.InsertOne(ctx, menu)
		rt, _ := ret.InsertOne(ctx, menu)
		c.JSON(http.StatusOK, rt)
	}
}

func GetMenu() gin.HandlerFunc {

	return func(c *gin.Context) {

		var menu models.Menu

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		menuId := c.Param("menu_id")

		err := ret.FindOne(ctx, bson.M{"food_id": menuId}).Decode(&menu)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, menu)
	}
}

func UpdateMenu() gin.HandlerFunc {

	return func(c *gin.Context) {

		var menu models.Menu

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		menuId := c.Param("Menu_id")

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D

		updateObj = append(updateObj, bson.E{"name", menu.Name})

		//updateObj = append(updateObj, bson.E{"name", food.Food_image})

		//updateObj = append(updateObj, bson.E{"name", food.Price})

		upsert := true

		ops := options.UpdateOptions{
			Upsert: &upsert,
		}

		filter := bson.M{"menu_id": menuId}

		result, err := ret.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &ops)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, result)
	}
}

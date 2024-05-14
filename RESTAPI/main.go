package main

import (
	user "User/User"
	"User/auth"
	"User/campaign"
	"User/handler"
	"User/helper"
	"User/payment"
	"User/transaction"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {

	dsn := "root:Vijay@123@tcp(localhost:3306)/appdb?charset=utf8mb4&parseTime=True&loc=Local" // %40
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db) // created struct object

	campaignRepository := campaign.NewRepository(db) // created struct object

	transactionRepository := transaction.NewRepository(db) // created struct object

	userService := user.NewService(userRepository) // created interface object

	campaignService := campaign.NewService(campaignRepository) // created interface object

	authService := auth.Newservice()

	paymentService := payment.NewService()

	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)         // passing the interface object
	campaignHandler := handler.NewCampaignHandler(campaignService)          // passing the interface object
	transactionHandler := handler.NewTransactionHandler(transactionService) // passing the interface object

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, strict-origin-when-cross-origin, Authorization"},
	}))

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("golangstarter", cookieStore))

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)                                             // register api working
	api.POST("/sessions", userHandler.Login)                                                 // Login  working
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)                          // email check working
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar) //working
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser) // working

	api.GET("/campaigns", campaignHandler.GetCampaigns)                                              // working
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)                                           // not filtering
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign) // working
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	//put api not filtering correctly
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	// create transactions- database has payment url column and input json format not defined.
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

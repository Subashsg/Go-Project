package main

import (
	"employeeattenance/dbconnection"
	"employeeattenance/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := dbconnection.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	router.POST("/checkin", handlers.CheckIn)
	router.PUT("/checkout/:id", handlers.CheckOut)
	router.GET("/report/:id", handlers.AttendanceReport)
	router.GET("/attendance", handlers.FilteredAttendance)

	router.Run(":8080")
}

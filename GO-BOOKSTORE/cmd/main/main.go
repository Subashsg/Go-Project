package main

import (
	"log"
	"net/http"

	"GO-BOOKSTORE/pkg/routes"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := gin.Default()          // Use gin.Default() instead of gin.NewRouter()
	routes.RegisterBookStore(r) // Corrected function name
	http.Handle("/", r)
	// Use r directly with http.ListenAndServe
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

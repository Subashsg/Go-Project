package dbconnection

import (
	"log"

	"employeeattenance/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
 
func ConnectDB() error {
	var err error

	db, err = gorm.Open("mysql", "root:Subashsvs@85@tcp(127.0.0.1:3306)/attendance?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database successfully")

	if err := db.AutoMigrate(&models.Employee{}).Error; err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

func GetDB() *gorm.DB {
	return db
}

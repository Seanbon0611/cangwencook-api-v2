package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// "os"
	// "github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// db, err := gorm.Open("postgres", "host=localhost port=5431 user=postgres dbname=test_db sslmode=disable")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Failed to connect to Database")
	// }
	// fmt.Println("Database Connected")
	// defer db.Close()
	r.Run()
}

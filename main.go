package main

import (
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//function that is able to get data from .env file
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	return os.Getenv(key)
}
func main() {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	// db, err := gorm.Open("postgres", "host=localhost port=5431 user=postgres dbname=test_db sslmode=disable")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Failed to connect to Database")
	// }
	// fmt.Println("Database Connected")
	// defer db.Close()

	//starts server
	server.Run(":3001")
}

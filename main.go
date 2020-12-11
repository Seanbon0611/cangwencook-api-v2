package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()
	// db, err := gorm.Open("postgres", "host=localhost port=5431 user=postgres dbname=test_db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Failed to connect to Database")
	// }
	// fmt.Println("Database Connected")
	// defer db.Close()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.Run()
}

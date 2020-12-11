package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5431 user=postgres dbname=test_db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

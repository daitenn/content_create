package main

import (
	"fmt"
	"go-restapi/db"
	"go-restapi/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Content{})
}

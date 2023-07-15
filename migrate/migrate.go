package main

import (
	"fmt"
	"shift-app-go-api/db"
	"shift-app-go-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Workspace{}, &model.Shift{}, &model.HourlyWage{}, &model.Incom{})
}

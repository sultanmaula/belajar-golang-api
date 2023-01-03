package main

import (
	"belajar-golang-api/handler"
	"belajar-golang-api/rooms"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:sultan123@tcp(127.0.0.1:3306)/belajar-golang-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection errors!")
	}

	db.AutoMigrate(&rooms.Room{})

	// CRUD
	// ===============
	// CREATE DATA
	// ===============

	// rooms := rooms.Room{}
	// rooms.Title = "Economy Room"
	// rooms.Description = "Kamar dengan kelas ekonomi"
	// rooms.Price = 125000
	// rooms.Rating = 5

	// err = db.Create(&rooms).Error
	// if err != nil {
	// 	fmt.Println("Error creating rooms record!")
	// }

	var rooms []rooms.Room

	err = db.Debug().Find(&rooms).Error
	if err != nil {
		fmt.Println("Error finding rooms record!")
	}

	for _, b := range rooms {
		fmt.Println("Title :", b.Title)
		fmt.Println("Rooms object %v", b)
	}

	router := gin.Default()
	router.GET("/", handler.RootHandler)
	router.GET("/test", handler.TestHandler)
	router.GET("/rooms/:price/:title", handler.RoomsHandler)
	router.GET("/menus", handler.MenusHandler)
	router.POST("/rooms", handler.PostRoomsHandler)
	router.Run()
}

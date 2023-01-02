package main

import (
	"belajar-golang-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handler.RootHandler)
	router.GET("/test", handler.TestHandler)
	router.GET("/rooms/:price/:title", handler.RoomsHandler)
	router.GET("/menus", handler.MenusHandler)
	router.POST("/rooms", handler.PostRoomsHandler)
	router.Run()
}

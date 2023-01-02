package handler

import (
	"belajar-golang-api/rooms"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Sultan Maula",
		"bio":  "Web Developer",
	})
}

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Test",
	})
}

func RoomsHandler(c *gin.Context) {
	price := c.Param("price")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"id":    price,
	})
}

func MenusHandler(c *gin.Context) {
	price := c.Query("price")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostRoomsHandler(c *gin.Context) {
	var roomInput rooms.RoomInput

	err := c.ShouldBindJSON(&roomInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": roomInput.Title,
		"price": roomInput.Price,
	})
}

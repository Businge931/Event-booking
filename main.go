package main

import (
	"net/http"
	"strconv"

	"example.com/event-booking/db"
	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.POST("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}


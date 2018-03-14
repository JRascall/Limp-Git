package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

var events []string

func main() {
	events = make([]string, 0)

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("/public", true)))

	router.GET("/events", func(c *gin.Context) {
		events = append(events[:0], events[1:]...)
		c.JSON(200, events)
	})

	router.POST("/hook", func(c *gin.Context){
		eventType := c.GetHeader("X-GitHub-Event")
		events = append(events, eventType)
		c.Status(200)
	})

	router.Run(":8005")
}

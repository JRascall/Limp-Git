package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"fmt"
)

var events []string

func main() {
	events = make([]string, 0)

	router := gin.Default()
	//router.LoadHTMLGlob(("/templates/*"))
	router.Use(static.Serve("/", static.LocalFile("/public", true)))

	router.GET("/events", func(c *gin.Context) {

	})

	router.POST("/hook", func(c *gin.Context){
		eventType := c.GetHeader("X-GitHub-Event")
		events = append(events, eventType)
		fmt.Println(events)
		c.Status(200)
	})

	router.Run(":8005")
}

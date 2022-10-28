package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Info struct {
	SlackUsername string `json:"slack_username"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", getting)
	router.Run(":8080")
}

func getting(ctx *gin.Context) {
	details := Info{
		SlackUsername: "Adeben33",
		Backend:       true,
		Age:           25,
		Bio:           "I am a beginner golang developer interest in creating worthy portfolio project",
	}
	ctx.JSON(200, details)
}

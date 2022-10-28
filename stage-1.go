package main

import (
	"github.com/gin-gonic/gin"
)

type Info struct {
	SlackUsername string `json:"slack_username"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	router := gin.Default()
	//router.Use(cors.Default())
	router.Use(CORSMiddleware())
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

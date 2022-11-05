package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Info struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}
type Operation struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
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
	router.POST("/operation", enum)
	router.Run(":8080")
}
func enum(c *gin.Context) {
	var newOperation Operation
	if err := c.BindJSON(&newOperation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	type returnStruct struct {
		SlackUsername string `json:"slackUsername"`
		Result        int    `json:"result"`
		OperationType string `json:"operation_type"`
	}

	if strings.Index(newOperation.OperationType, "add") != -1 {
		re := regexp.MustCompile("[0-9]+")
		numberSlice := re.FindAllString(string(newOperation.OperationType), -1)
		if re.MatchString(string(newOperation.OperationType)) {
			total := 0
			for _, value := range numberSlice {
				value, _ := strconv.Atoi(value)
				total += value
			}
			var addedNumberReturnStruct returnStruct
			//addedNumber := newOperation.X + newOperation.Y
			addedNumberReturnStruct.SlackUsername = "Adeben33"
			addedNumberReturnStruct.Result = total
			addedNumberReturnStruct.OperationType = newOperation.OperationType
			addedNumberReturnStruct.OperationType = "addition"
			c.JSON(200, addedNumberReturnStruct)
		} else {
			var addedNumberReturnStruct returnStruct
			addedNumber := newOperation.X + newOperation.Y
			addedNumberReturnStruct.SlackUsername = "Adeben33"
			addedNumberReturnStruct.Result = addedNumber
			addedNumberReturnStruct.OperationType = newOperation.OperationType
			addedNumberReturnStruct.OperationType = "addition"
			c.JSON(200, addedNumberReturnStruct)
		}
	} else if strings.Index(newOperation.OperationType, "sub") != -1 {
		re := regexp.MustCompile("[0-9]+")
		numberSlice := re.FindAllString(string(newOperation.OperationType), -1)
		if re.MatchString(string(newOperation.OperationType)) {
			//total := 0
			//for _, value := range numberSlice {
			//	value, _ := strconv.Atoi(value)
			//	total = value - total
			//}
			var addedNumberReturnStruct returnStruct
			Number1, _ := strconv.Atoi(numberSlice[0])
			Number2, _ := strconv.Atoi(numberSlice[1])
			addedNumberReturnStruct.SlackUsername = "Adeben33"
			addedNumberReturnStruct.Result = Number1 - Number2
			addedNumberReturnStruct.OperationType = newOperation.OperationType
			addedNumberReturnStruct.OperationType = "subtraction"
			c.JSON(200, addedNumberReturnStruct)
		} else {
			var NumberReturnStruct returnStruct
			Number := newOperation.X - newOperation.Y
			NumberReturnStruct.SlackUsername = "Adeben33"
			NumberReturnStruct.Result = Number
			NumberReturnStruct.OperationType = newOperation.OperationType
			NumberReturnStruct.OperationType = "subtraction"
			c.JSON(200, NumberReturnStruct)
		}
	} else if strings.Index(newOperation.OperationType, "mul") != -1 {
		re := regexp.MustCompile("[0-9]+")
		numberSlice := re.FindAllString(string(newOperation.OperationType), -1)
		if re.MatchString(string(newOperation.OperationType)) {
			total := 1
			for _, value := range numberSlice {
				value, _ := strconv.Atoi(value)
				total *= value
			}
			var addedNumberReturnStruct returnStruct
			//addedNumber := newOperation.X + newOperation.Y
			addedNumberReturnStruct.SlackUsername = "Adeben33"
			addedNumberReturnStruct.Result = total
			addedNumberReturnStruct.OperationType = newOperation.OperationType
			addedNumberReturnStruct.OperationType = "multiplication"
			c.JSON(200, addedNumberReturnStruct)
		} else {
			var NumberReturnStruct returnStruct
			Number := newOperation.X * newOperation.Y
			NumberReturnStruct.SlackUsername = "Adeben33"
			NumberReturnStruct.Result = Number
			NumberReturnStruct.OperationType = newOperation.OperationType
			NumberReturnStruct.OperationType = "multiplication"
			c.JSON(200, NumberReturnStruct)
		}
	}
}

func getting(c *gin.Context) {
	details := Info{
		SlackUsername: "Adeben33",
		Backend:       true,
		Age:           25,
		Bio:           "I am a beginner golang developer interest in creating worthy portfolio project",
	}
	c.PureJSON(200, details)
}

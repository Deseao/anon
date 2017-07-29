package main

import (
	"fmt"
	"github.com/Deseao/anon/api/internal/group"
	"github.com/Deseao/anon/api/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	groupHandler := group.GroupHandler{}
	router.Use(middleware.GroupHandler(&groupHandler))
	router.GET("/ping", Ping)
	router.POST("/create", CreateGroup)
	router.POST("/signup", Signup)
	router.POST("/send", SendMessage)
	router.Run(":8080")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func CreateGroup(c *gin.Context) {
	groupHandler := c.MustGet("groupHandler").(*group.GroupHandler)
	groupCode := groupHandler.Create()
	fmt.Println("Group Created: ", groupHandler.Groups)
	response := NewGroupPayload{Code: groupCode}
	c.JSON(http.StatusOK, response)
}

func Signup(c *gin.Context) {
	groupHandler := c.MustGet("groupHandler").(*group.GroupHandler)
	var signupPayload SignupPayload
	c.BindJSON(&signupPayload)
	fmt.Println("Signup Payload: ", signupPayload)
	groupExists := groupHandler.GroupExists(signupPayload.Code)
	fmt.Println("Group Exists", groupExists)
	if !groupExists {
		c.Status(http.StatusBadRequest)
		return
	}
	err := groupHandler.AddParticipant(signupPayload.Code, signupPayload.Email, signupPayload.Phone)
	if err != nil {
		fmt.Println("Error from adding participant: ", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
	return
}

func SendMessage(c *gin.Context) {
	c.Status(http.StatusOK)
}

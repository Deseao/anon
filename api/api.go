package main

import (
	"github.com/Deseao/anon/api/internal/group"
	"github.com/Deseao/anon/api/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(middleware.GroupHandler)
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
	groupHandler := c.MustGet("groupHandler").(group.GroupHandler)
	groupCode := groupHandler.Create()
	response := NewGroupPayload{Code: groupCode}
	c.JSON(http.StatusOK, response)
}

func Signup(c *gin.Context) {
	c.Status(http.StatusOK)
}

func SendMessage(c *gin.Context) {
	c.Status(http.StatusOK)
}

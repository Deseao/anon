package main

import (
	"fmt"
	"github.com/Deseao/anon/internal/config"
	"github.com/Deseao/anon/internal/group"
	"github.com/Deseao/anon/internal/middleware"
	"github.com/Deseao/anon/internal/participant"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	conf, err := config.ReadConfig("api-config.toml")
	if err != nil {
		log.Fatal("Could not read config: ", err)
	}
	participant.Email_api_key = conf.SendGrid.Key
	participant.From_address = conf.SendGrid.FromAddress
	participant.Twilio_sid = conf.Twilio.SID
	participant.Twilio_api_key = conf.Twilio.Key
	participant.Twilio_account_id = conf.Twilio.AccountID
	participant.Twilio_from_number = conf.Twilio.FromNumber
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
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
	groupHandler := c.MustGet("groupHandler").(*group.GroupHandler)
	var messagePayload MessagePayload
	c.BindJSON(&messagePayload)
	fmt.Println("Message Payload", messagePayload)
	err := groupHandler.SendMessage(messagePayload.Code, messagePayload.Message)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

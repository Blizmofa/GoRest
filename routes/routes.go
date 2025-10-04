package routes

import (
	"example.com/go-rest/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	pathEvents          = "/events"
	pathEventByID       = "/events/:id"
	pathRegisterByEvent = "/events/:id/register"
	pathSignup          = "/signup"
	pathLogin           = "/login"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET(pathEvents, getEvents)
	server.GET(pathEventByID, getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST(pathEvents, createEvent)
	authenticated.PUT(pathEventByID, updateEvent)
	authenticated.DELETE(pathEventByID, deleteEvent)
	authenticated.POST(pathRegisterByEvent, registerForEvent)
	authenticated.DELETE(pathRegisterByEvent, cancelRegistration)

	server.POST(pathSignup, signUp)
	server.POST(pathLogin, login)
}

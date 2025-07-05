package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/events", app.getAllEvents)
		v1.GET("/event/:id", app.getEvent)

		v1.GET("/event/:id/attendees", app.getAttendeesForEvent)
		v1.GET("/attendees/:id/events", app.getEventsByAttendee)

		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.login)
	}

	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleware())
	{
		authGroup.POST("/event", app.createEvent)
		authGroup.PUT("/event/:id", app.updateEvent)
		authGroup.DELETE("/event/:id", app.deleteEvent)

		authGroup.POST("/event/:id/attendee/:userId", app.addAttendeeToEvent)
		authGroup.DELETE("/events/:id/attendees/:userId", app.deleteAttendeeFromEvent)
	}

	return g
}

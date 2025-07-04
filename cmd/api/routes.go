package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.POST("/event", app.createEvent)
		v1.GET("/events", app.getAllEvents)
		v1.GET("/event/:id", app.getEvent)
		v1.PUT("/event/:id", app.updateEvent)
		v1.DELETE("/event/:id", app.deleteEvent)
		v1.POST("/event/:id/attendee/:userId", app.addAttendeeToEvent)
		v1.GET("/event/:id/attendees", app.getAttendeesForEvent)

		v1.POST("/auth/register", app.registerUser)
	}

	return g
}

package main

import (
	"event-app/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.Insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create event",
		})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve events",
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID",
		})
		return
	}

	event, err := app.models.Events.Get(id)

	if event != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve event",
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event id",
		})
		return
	}

	existingEvent, err := app.models.Events.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve event",
		})
		return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

	updatedEvent := &database.Event{}

	if err := c.ShouldBindJSON(updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedEvent.Id = id

	if err := app.models.Events.Update(updatedEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update event",
		})
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event Id",
		})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete event",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (app *application) addAttendeeToEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event id",
		})
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	event, err := app.models.Events.Get(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

	userToAdd, err := app.models.Users.Get(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user",
		})
		return
	}

	if userToAdd == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	existingAttendee, err := app.models.Attendees.GetByEventAndAttendee(event.Id, userToAdd)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve attendeee",
		})
		return
	}

	if existingAttendee != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Attendee already exists",
		})
		return
	}

	attendee := database.Attendee{
		EventId: event.Id,
		UserId:  userToAdd.Id,
	}

	_, err = app.models.Attendees.Insert(&attendee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add attendeee",
		})
		return
	}

	c.JSON(http.StatusCreated, attendee)
}

func (app *application) getAttendeesForEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event id",
		})
		return
	}

	users, err := app.models.Attendees.GetAttendeesByEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve attendees for events",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

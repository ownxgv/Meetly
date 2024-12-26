package router

import (
	"net/http"

	"binary_horizon/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", postEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)
	return r
}

func postEvent(ctx *gin.Context) {
	var event db.Event
	err := ctx.Bind(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateEvent(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"event": res,
	})
}

func getEvents(ctx *gin.Context) {
	res, err := db.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"events": res,
	})
}

func getEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetEvent(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"event": res,
	})
}

func updateEvent(ctx *gin.Context) {
	var updatedEvent db.Event
	err := ctx.Bind(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbEvent, err := db.GetEvent(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbEvent.Name = updatedEvent.Name
	dbEvent.Description = updatedEvent.Description

	res, err := db.UpdateEvent(dbEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"task": res,
	})
}

func deleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteEvent(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}

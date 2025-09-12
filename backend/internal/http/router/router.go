// Package router provides a router for the application
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pollparlor/internal/http/handler"
)

// New creates a new router
func New(h *handler.PollHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello, world!"}) })

	api := r.Group("/api")
	api.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	v1 := api.Group("/v1")

	// Polls
	{
		polls := v1.Group("/polls")
		polls.GET("", h.List)
		polls.GET("/:id", h.Get)
		polls.POST("", h.Create)
	}
	return r
}

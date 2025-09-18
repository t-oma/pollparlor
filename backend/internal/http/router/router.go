// Package router provides a router for the application
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pollparlor/internal/http/handler"
	mw "pollparlor/internal/http/middleware"
)

// New creates a new router
func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(mw.RequestLogger())

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello, world!"}) })

	return r
}

// CreateHealthRoutes creates a new health routes
func CreateHealthRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })

	return r
}

// CreatePollRoutes creates a new poll routes
func CreatePollRoutes(r *gin.RouterGroup, h *handler.PollHandler) *gin.RouterGroup {
	polls := r.Group("/polls")
	polls.GET("", h.List)                        // all polls
	polls.GET("/:id", h.Get)                     // get poll by id
	polls.POST("", h.Create)                     // create poll
	polls.DELETE("/:id", h.Delete)               // delete poll
	polls.POST("/with-items", h.CreateWithItems) // create poll with items
	polls.GET("/:id/items", h.GetItems)          // get poll items
	polls.POST("/:id/items", h.CreateItems)      // create poll items
	polls.GET("/:id/pairs", h.GetPairs)          // get poll pairs

	return polls
}

// CreateUserRoutes creates a new user routes
func CreateUserRoutes(r *gin.RouterGroup, h *handler.UserHandler) *gin.RouterGroup {
	users := r.Group("/users")
	users.GET("", h.List)    // all users
	users.GET("/:id", h.Get) // get user by id
	users.POST("", h.Create) // create user

	return users
}

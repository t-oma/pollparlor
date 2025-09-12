// Package server provides a server for the application
package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Server is a server
type Server struct {
	cfg    *any
	router *gin.Engine
}

// NewServer creates a new server
func NewServer(cfg *any) *Server {
	srv := &Server{
		cfg:    cfg,
		router: gin.Default(),
	}

	srv.initRoutes()

	return srv
}

type UserVerification struct {
	Valid bool `json:"valid"`
}

type User struct {
	UUID              string           `json:"uuid"`
	Email             string           `json:"email"`
	CreatedAt         time.Time        `json:"createdAt"`
	PasswordChangedAt time.Time        `json:"passwordChangedAt"`
	Verification      UserVerification `json:"verification"`
}

type Poll struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Author    User      `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var polls = []Poll{
	{
		UUID:  uuid.NewString(),
		Title: "Programming languages",
		Author: User{
			UUID:              uuid.NewString(),
			Email:             "user1@example.com",
			CreatedAt:         time.Now(),
			PasswordChangedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		UUID:  uuid.NewString(),
		Title: "Web frameworks",
		Author: User{
			UUID:              uuid.NewString(),
			Email:             "user2@example.com",
			CreatedAt:         time.Now(),
			PasswordChangedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(time.Hour * 2),
	},
}

type GetPollRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

func (s *Server) initRoutes() {
	s.router.Use(cors.Default())

	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	s.router.GET("/polls", func(c *gin.Context) {
		c.JSON(200, polls)
	})
	s.router.GET("/polls/:uuid", func(c *gin.Context) {
		var req GetPollRequest

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, poll := range polls {
			if poll.UUID == req.UUID {
				c.JSON(200, poll)
				return
			}
		}

		c.JSON(404, gin.H{
			"error": "poll not found",
		})
	})
}

// Run runs the server
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

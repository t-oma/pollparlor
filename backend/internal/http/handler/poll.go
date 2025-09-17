// Package handler provides handlers for the application
package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"pollparlor/internal/domain"
	"pollparlor/internal/service/poll"
)

// PollHandler is a handler for polls
type PollHandler struct {
	svc *poll.Service
}

// NewPollHandler creates a new handler for polls
func NewPollHandler(s *poll.Service) *PollHandler { return &PollHandler{svc: s} }

type listPollURI struct {
	Limit int64 `form:"limit" binding:"omitempty,min=1,max=100"`
	Skip  int64 `form:"skip" binding:"omitempty,min=0"`
}

// List returns a list of polls
func (h *PollHandler) List(c *gin.Context) {
	var query listPollURI
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := h.svc.List(query.Limit, query.Skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if items == nil {
		c.JSON(http.StatusOK, []domain.Poll{})
		return
	}
	c.JSON(http.StatusOK, items)
}

type getPollURI struct {
	ID bson.ObjectID `uri:"id" binding:"required,uuid"`
}

// Get returns a poll by ID
func (h *PollHandler) Get(c *gin.Context) {
	var uri getPollURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := h.svc.Get(uri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "poll not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

type createPollJSON struct {
	Title string `json:"title" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
}

// Create creates a new poll
func (h *PollHandler) Create(c *gin.Context) {
	var body createPollJSON
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	p := domain.Poll{
		ID:    bson.NewObjectID(),
		Title: body.Title,
		Author: domain.User{
			UUID:              uuid.NewString(),
			Email:             body.Email,
			CreatedAt:         now,
			PasswordChangedAt: now,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := h.svc.Create(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

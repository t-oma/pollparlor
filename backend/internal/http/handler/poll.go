// Package handler provides handlers for the application
package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"pollparlor/internal/domain"
	"pollparlor/internal/service/poll"
)

type PollHandler struct {
	svc *poll.Service
}

func NewPollHandler(s *poll.Service) *PollHandler { return &PollHandler{svc: s} }

type listPollURI struct {
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
	Skip  int `form:"skip" binding:"omitempty,min=0"`
}

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
	ID string `uri:"id" binding:"required,uuid"`
}

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

func (h *PollHandler) Create(c *gin.Context) {
	var body createPollJSON
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	p := domain.Poll{
		ID:    uuid.NewString(),
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

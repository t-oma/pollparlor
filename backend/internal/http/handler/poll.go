// Package handler provides handlers for the application
package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	Title    string `json:"title" binding:"required,min=3"`
	AuthorID string `json:"authorId" binding:"required,uuid"`
}

// Create creates a new poll
func (h *PollHandler) Create(c *gin.Context) {
	var body createPollJSON
	if err := c.ShouldBindJSON(&body); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty body"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	p := domain.Poll{
		ID:        bson.NewObjectID(),
		Title:     body.Title,
		AuthorID:  body.AuthorID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := h.svc.Create(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

type deletePollURI struct {
	ID string `uri:"id" binding:"required"`
}

// Delete deletes a poll
func (h *PollHandler) Delete(c *gin.Context) {
	var uri deletePollURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oid, err := bson.ObjectIDFromHex(uri.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := h.svc.Get(oid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "poll not found"})
		return
	}
	if err := h.svc.Delete(p.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "poll with id " + uri.ID + " deleted"})
}

type getPollItemsURI struct {
	PollID string `uri:"id" binding:"required"`
}

// GetItems returns a list of items in the poll
func (h *PollHandler) GetItems(c *gin.Context) {
	var uri getPollItemsURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oid, err := bson.ObjectIDFromHex(uri.PollID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items, err := h.svc.GetItems(oid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "poll not found"})
		return
	}
	c.JSON(http.StatusOK, items)
}

type createPollItemsJSON struct {
	PollID bson.ObjectID      `json:"pollId" binding:"required"`
	Items  []domain.Candidate `json:"items" binding:"required,min=4,max=100"`
}

// CreateItems creates a new list of items in the poll
func (h *PollHandler) CreateItems(c *gin.Context) {
	var body createPollItemsJSON
	if err := c.ShouldBindJSON(&body); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty body"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items := body.Items
	if len(items) < 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number of candidates. Provide at least 4"})
		return
	}
	if len(items)&1 != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number of candidates. Should be even"})
		return
	}
	for i := range items {
		items[i].ID = bson.NewObjectID()
	}

	err := h.svc.CreateItems(body.PollID, items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, items)
}

type getPollPairsURI struct {
	PollID bson.ObjectID `uri:"pollId" binding:"required"`
}

// GetPairs returns a list of pairs in the poll
func (h *PollHandler) GetPairs(c *gin.Context) {
	var uri getPollPairsURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pairs, err := h.svc.GetPairs(uri.PollID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "poll not found"})
		return
	}
	c.JSON(http.StatusOK, pairs)
}

type createPollWithItemsJSON struct {
	createPollJSON
	Items []domain.Candidate `json:"items" binding:"required,min=4,max=100"`
}

// CreateWithItems creates a new poll with items
func (h *PollHandler) CreateWithItems(c *gin.Context) {
	var body createPollWithItemsJSON
	if err := c.ShouldBindJSON(&body); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty body"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items := body.Items
	if len(items) < 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number of candidates. Provide at least 4"})
		return
	}
	if len(items)&1 != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number of candidates. Should be even"})
		return
	}

	p := domain.Poll{
		ID:        bson.NewObjectID(),
		Title:     body.Title,
		AuthorID:  body.AuthorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := h.svc.Create(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range items {
		items[i].ID = bson.NewObjectID()
	}

	err := h.svc.CreateItems(p.ID, items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"poll": p, "items": items})
}

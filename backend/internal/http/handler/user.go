package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"pollparlor/internal/domain"
	"pollparlor/internal/service/user"
)

// UserHandler is a handler for users
type UserHandler struct {
	svc *user.Service
}

// NewUserHandler creates a new handler for users
func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{svc: s}
}

type listUserURI struct {
	Limit int64 `form:"limit" binding:"omitempty,min=1,max=100"`
	Skip  int64 `form:"skip" binding:"omitempty,min=0"`
}

// List returns a list of users
func (h *UserHandler) List(c *gin.Context) {
	var query listUserURI
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
		c.JSON(http.StatusOK, []domain.User{})
		return
	}
	c.JSON(http.StatusOK, items)
}

type getUserURI struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// Get returns a user by ID
func (h *UserHandler) Get(c *gin.Context) {
	var uri getUserURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.svc.Get(uri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

type createUserJSON struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Create creates a new user
func (h *UserHandler) Create(c *gin.Context) {
	var body createUserJSON
	if err := c.ShouldBindJSON(&body); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty body"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	user := domain.User{
		UUID:              uuid.NewString(),
		Email:             body.Email,
		Password:          body.Password,
		CreatedAt:         now,
		PasswordChangedAt: now,
		Verification: domain.UserVerification{
			Valid: false,
		},
	}
	if err := h.svc.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

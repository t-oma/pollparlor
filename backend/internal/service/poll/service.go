// Package poll provides services for the application
package poll

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"pollparlor/internal/domain"
)

// Service for polls
type Service struct {
	repo domain.PollRepository
}

// New creates a new polls service
func New(repo domain.PollRepository) *Service {
	return &Service{repo: repo}
}

// List returns a list of polls
func (s *Service) List(limit, skip int64) ([]domain.Poll, error) {
	return s.repo.List(limit, skip)
}

// Get returns a poll by ID
func (s *Service) Get(id bson.ObjectID) (*domain.Poll, error) {
	return s.repo.GetByID(id)
}

// Create creates a new poll
func (s *Service) Create(p domain.Poll) error {
	return s.repo.Create(p)
}

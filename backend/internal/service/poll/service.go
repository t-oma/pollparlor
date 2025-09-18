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
func (s *Service) List(limit, skip int64) ([]domain.PollWithAuthor, error) {
	return s.repo.List(limit, skip)
}

// Get returns a poll by ID
func (s *Service) Get(id bson.ObjectID) (*domain.PollWithAuthor, error) {
	return s.repo.GetByID(id)
}

// Create creates a new poll
func (s *Service) Create(p domain.Poll) error {
	return s.repo.Create(p)
}

// Delete deletes a poll
func (s *Service) Delete(id bson.ObjectID) error {
	return s.repo.Delete(id)
}

// GetItems returns a list of items
func (s *Service) GetItems(pollID bson.ObjectID) ([]domain.Candidate, error) {
	return s.repo.GetItems(pollID)
}

// CreateItems creates a new list of items
func (s *Service) CreateItems(pollID bson.ObjectID, items []domain.Candidate) error {
	return s.repo.CreateItems(pollID, items)
}

// GetPairs returns a list of pairs
func (s *Service) GetPairs(pollID bson.ObjectID) ([]domain.PollPair, error) {
	return s.repo.GetPairs(pollID)
}

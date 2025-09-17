// Package user provides users services for the application
package user

import "pollparlor/internal/domain"

// Service for users
type Service struct {
	repo domain.UserRepository
}

// New creates a new user service
func New(repo domain.UserRepository) *Service {
	return &Service{repo: repo}
}

// List returns a list of users
func (s *Service) List(limit, skip int64) ([]domain.User, error) {
	return s.repo.List(limit, skip)
}

// Get returns a user by ID
func (s *Service) Get(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

// Create creates a new user
func (s *Service) Create(u domain.User) error {
	return s.repo.Create(u)
}

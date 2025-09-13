// Package domain provides domain objects like structs and interfaces
package domain

import "time"

type Poll struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    User      `json:"author"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PollRepository interface {
	List(limit, skip int) ([]Poll, error)
	GetByID(id string) (*Poll, error)
	Create(p Poll) error
}

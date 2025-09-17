// Package domain provides domain objects like structs and interfaces
package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Candidate represents a candidate in the poll
type Candidate struct {
	ID    bson.ObjectID `bson:"_id"   json:"id"`
	Title string        `bson:"title" json:"title"`
}

// PollItems represents items in the poll
type PollItems struct {
	ID     bson.ObjectID `bson:"_id"    json:"id"`
	PollID bson.ObjectID `bson:"pollId" json:"pollId"`
	Cands  []Candidate   `bson:"cands"  json:"cands"`
}

// PollPair represents a pair of candidates in the poll
type PollPair struct {
	Left  Candidate `bson:"left"  json:"left"`
	Right Candidate `bson:"right" json:"right"`
}

// Poll represents a poll in the system
type Poll struct {
	ID        bson.ObjectID `bson:"_id"       json:"id"`
	Title     string        `bson:"title"     json:"title"`
	Author    User          `bson:"author"    json:"author"`
	Likes     int           `bson:"likes"     json:"likes"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

// PollRepository is a repository for polls
type PollRepository interface {
	List(limit, skip int64) ([]Poll, error)
	GetByID(id bson.ObjectID) (*Poll, error)
	Create(p Poll) error
	GetItems(pollID bson.ObjectID) ([]Candidate, error)
	CreateItems(pollID bson.ObjectID, items []Candidate) error
	GetPairs(pollID bson.ObjectID) ([]PollPair, error)
}

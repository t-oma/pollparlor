package domain

import "time"

// UserVerification represents a user verification status
type UserVerification struct {
	Valid bool `bson:"valid" json:"valid"`
}

// User represents a user in the system
type User struct {
	UUID              string           `bson:"uuid"              json:"uuid"`
	Email             string           `bson:"email"             json:"email"`
	Password          string           `bson:"password"          json:"password"`
	CreatedAt         time.Time        `bson:"createdAt"         json:"createdAt"`
	PasswordChangedAt time.Time        `bson:"passwordChangedAt" json:"passwordChangedAt"`
	Verification      UserVerification `bson:"verification"      json:"verification"`
}

// UserRepository is a repository for users
type UserRepository interface {
	List(limit, skip int64) ([]User, error)
	GetByID(id string) (*User, error)
	Create(p User) error
}

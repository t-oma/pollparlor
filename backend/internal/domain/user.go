package domain

import "time"

// UserVerification represents a user verification status
type UserVerification struct {
	Valid bool `json:"valid"`
}

// User represents a user in the system
type User struct {
	UUID              string           `json:"uuid"`
	Email             string           `json:"email"`
	CreatedAt         time.Time        `json:"createdAt"`
	PasswordChangedAt time.Time        `json:"passwordChangedAt"`
	Verification      UserVerification `json:"verification"`
}

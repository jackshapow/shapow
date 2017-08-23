package model

import (
	"time"
)

// PrimaryID is an ID field common to most models
type PrimaryID struct {
	ID uint `json:"id"`
}

// CommonDates is a set of common dates fields
type CommonDates struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at"`
}

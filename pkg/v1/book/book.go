package book

import (
	"time"
)

type Book struct {
	Id             int64     `json:"id,omitempty"`
	Title          string    `json:"title" validate:"required,max=255"`
	Description    string    `json:"description" validate:"required,max=1000"`
	AgeGroup       int       `json:"age_group" validate:"required"`
	PublishingDate time.Time `json:"publishing_date" validate:"required"`
}

package domain

import "time"

type Book struct {
	Id             int64
	Title          string
	Description    string
	AgeGroup       int
	PublishingDate time.Time
}

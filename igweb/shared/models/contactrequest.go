package models

import "time"

type ContactRequest struct {
	FirstName       string
	LastName        string
	Email           string
	Message         string
	TimeSensitivity time.Time
}

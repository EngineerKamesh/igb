package models

import "time"

type Gopher struct {
	Name      string
	Title     string
	Biodata   string
	ImageURI  string
	StartTime time.Time
}

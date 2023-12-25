package models

import "time"

type Families struct {
	ID          int8
	Name        string
	Kanji       string
	Description string
	Romaji      string
	Dub         string
	Level       int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

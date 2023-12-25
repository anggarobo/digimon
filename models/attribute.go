package models

import "time"

type Attribute struct {
	ID          int8
	Name        string
	Symbol      string
	Kanji       string
	Description string
	Romaji      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

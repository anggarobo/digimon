package models

import "time"

type Levels struct {
	ID          int8      `json:"id"`
	Name        string    `json:"name"`
	Kanji       string    `json:"kanji"`
	Description string    `json:"description"`
	Romaji      string    `json:"romaji,omitempty"`
	Level       *uint8    `json:"level"`
	Dub         string    `json:"dub,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

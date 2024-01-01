package models

import "time"

type Families struct {
	ID          int8      `json:"id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol,omitempty"`
	Kanji       string    `json:"kanji"`
	Description string    `json:"description,omitempty"`
	Romaji      string    `json:"romaji"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

package models

import "time"

type Levels struct {
	ID          int8      `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Kanji       string    `db:"kanji" json:"kanji"`
	Description string    `db:"description" json:"description"`
	Romaji      string    `db:"romaji" json:"romaji"`
	Level       int       `db:"level" json:"level"`
	Dub         string    `db:"dub" json:"dub"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

package models

type Digimon struct {
	ID string
	Name string
	Species // type
	Attribute
	Family []Family
	Group
	Attacks []Attacks
	Equipment string
	// Images []string
	Level string
}
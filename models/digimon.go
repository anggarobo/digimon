package models

type Digimon struct {
	ID string
	Name string
	Species // type
	Attribute
	Family []Family // field
	Group
	Attacks []Attacks
	Equipment string
	Images []string
	Level string
}

func (d Digimon) SpeciesName() string {
	return d.Species.Name
}

func (d Digimon) familyName() []string {
	var f []string
	for _, v := range d.Family {
		f = append(f, v.Name)
	}

	return f
}
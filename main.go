package main

import (
	"digimon/models"
	"fmt"
)

func main()  {
	d := models.Digimon{}

	d.Name = "agumon"
	d.Level = "child"
	d.Species.Name = "reptile"
	d.Attribute.Name = "vaccine"

	fmt.Println(d.SpeciesName())
}
package main

import (
	"digimon/models"
	"fmt"
)

func main()  {
	d := models.Digimon{}

	d.Name = "agumon"
	d.Level = "Child"
	d.Species.Name = "Reptile"
	d.Attribute.Name = "Vaccine"

	fmt.Println(d)
}
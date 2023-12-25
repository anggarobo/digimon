package main

import (
	"digimon/config"
	"fmt"
)

func main()  {
	fmt.Println("Digimons")
	
	db := config.Database{}

	db.Connect()
}
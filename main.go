package main

import (
	"fmt"
	"net/http"

	"digimon/config"
	h "digimon/handler"
)

var api string = "/api/v1"

func main()  {
	fmt.Println("Digimons")
	
	db := config.Database{}

	db.Connect()
	http.HandleFunc(api+"/levels", h.GetLevels)
	http.ListenAndServe(":9000", nil)
}
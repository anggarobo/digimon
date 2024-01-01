package main

import (
	"fmt"
	"net/http"

	"digimons/config"
	h "digimons/handler"
)

var api string = "/api/v1"

func main() {
	fmt.Println("Digimons")

	db := config.Database{}

	db.Connect()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Digimon"))
	})
	http.HandleFunc(api+"/attributes", h.GetAttributes)
	http.HandleFunc(api+"/families", h.GetFamilies)
	http.HandleFunc(api+"/levels", h.GetLevels)
	http.ListenAndServe(":9000", nil)
}

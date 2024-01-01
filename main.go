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
	http.HandleFunc(api+"/levels", h.GetLevels)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Digimon"))
	})
	http.ListenAndServe(":9000", nil)
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"digimons/config"
	"digimons/models"
)

func GetAttributes(w http.ResponseWriter, r *http.Request) {
	sql := config.SQL
	rows, err := sql.Query("SELECT id, name, kanji, description, romaji, symbol FROM attributes")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var sym []models.Attributes

	for rows.Next() {
		var each = models.Attributes{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Kanji, &each.Description, &each.Romaji, &each.Symbol)

		if err != nil {
			fmt.Println(err.Error())
		}

		sym = append(sym, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sym)
}

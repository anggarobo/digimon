package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"digimons/config"
	"digimons/models"
)

func GetFamilies(w http.ResponseWriter, r *http.Request) {
	sql := config.SQL
	rows, err := sql.Query("SELECT id, name, kanji, description, romaji, symbol FROM families")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var fam []models.Families

	for rows.Next() {
		var each = models.Families{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Kanji, &each.Description, &each.Romaji, &each.Symbol)

		if err != nil {
			fmt.Println(err.Error())
		}

		fam = append(fam, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fam)
}

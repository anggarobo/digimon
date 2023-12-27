package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"digimon/config"
	"digimon/models"
)

func GetLevels(w http.ResponseWriter, r *http.Request) {
	sql := config.SQL
	rows, err := sql.Query("SELECT * FROM levels")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var l []models.Levels

	for rows.Next() {
		var each = models.Levels{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Kanji, &each.Description, &each.Romaji, &each.Level, &each.Dub, &each.UpdatedAt, &each.CreatedAt)

		if err != nil {
			fmt.Println(err.Error() + "error")
		}

		l = append(l, each)
	}
	fmt.Println(l)

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(l)
}

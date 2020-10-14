package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/JakWai01/distributed-dungeons/pkg/models"
)

func GetAllCharactersHandler(db *sql.DB) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		dbCharacters := []models.Character{}

		rows, err := db.Query("select id, name,	player, occupation,	age, sex, residence, birthplace, strength, dexterity, power, constitution, appearance, education, size, intelligence, hitpoints, sanity, luck, magic_points, diceroll from characters")
		if err != nil {
			http.Error(rw, err.Error(), 500)

			return
		}

		for rows.Next() {
			var dbCharacter models.Character

			if err := rows.Scan(
				&dbCharacter.ID,
				&dbCharacter.Name,
				&dbCharacter.Player,
				&dbCharacter.Occupation,
				&dbCharacter.Age,
				&dbCharacter.Sex,
				&dbCharacter.Residence,
				&dbCharacter.Birthplace,
				&dbCharacter.Strength,
				&dbCharacter.Dexterity,
				&dbCharacter.Power,
				&dbCharacter.Constitution,
				&dbCharacter.Appearance,
				&dbCharacter.Education,
				&dbCharacter.Size,
				&dbCharacter.Intelligence,
				&dbCharacter.Hitpoints,
				&dbCharacter.Sanity,
				&dbCharacter.Luck,
				&dbCharacter.MagicPoints,
				&dbCharacter.Diceroll,
			); err != nil {
				http.Error(rw, err.Error(), 500)

				return
			}

			dbCharacters = append(dbCharacters, dbCharacter)
		}

		outCharacter, err := json.Marshal(dbCharacters)
		if err != nil {
			http.Error(rw, err.Error(), 500)

			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(outCharacter)
	}

	return http.HandlerFunc(fn)
}

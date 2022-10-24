package database

import (
	"database/sql"
	"edu/letu/wan/structs"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ClearPlayerTable() {
	db := OpenSQLite()
	defer db.Close()

	_, err := db.Exec(`
		DELETE FROM wan;`)
	if err != nil {
		log.Fatal(err)
	}
}

func playerExists(db *sql.DB, token, id string) bool {
	sqlStmt := `
		SELECT id FROM wan 
			WHERE token = ?
			OR id = ?;`
	err := db.QueryRow(sqlStmt, token, id).Scan()

	return err != sql.ErrNoRows
}

func AddPlayer(token string, player structs.PlayerInfo) bool {
	db := OpenSQLite()
	defer db.Close()

	if playerExists(db, token, player.ID) {
		return false
	}

	sqlStmt := `
		INSERT INTO wan (
			expires,
			token,
			id,
			name_adjective,
			name_noun,
			picture
		) VALUES (
			datetime('now', '1 day'),
			?, ?, ?, ?, ?
		);`
	_, err := db.Exec(sqlStmt, token, player.ID, player.Name.Adjective, player.Name.Noun, player.Picture)
	if err != nil {
		log.Fatal(err)
	}
	return true
}


type queryType string
const (
	byID queryType = "id"
	byToken queryType = "token"
)

// gets a player from the database WHERE query = value
// i.e. getPlayer("id", "1234") will return the player with id 1234
func getPlayer(query queryType, value string) *structs.PlayerInfo {
	db := OpenSQLite()
	defer db.Close()

	//sql statement
	sqlStmt := fmt.Sprintf(`
		SELECT id, name_adjective, name_noun, picture
			FROM wan WHERE %s = ?;
	`, query)

	//variables to store player info
	var id string
	var name_adjective, name_noun, picture int

	//query sql
	err := db.QueryRow(sqlStmt, value).Scan(&id, &name_adjective, &name_noun, &picture)
	if err == sql.ErrNoRows {
		return nil
	}

	//return player info
	return &structs.PlayerInfo{
		ID: id,
		Name: structs.PlayerName{
			Adjective: name_adjective,
			Noun: name_noun,
		},
		Picture: picture,
	}
}

func GetPlayerByID(id string) *structs.PlayerInfo {
	return getPlayer(byID, id)
}

func GetPlayerByToken(token string) *structs.PlayerInfo {
	return getPlayer(byToken, token)
}

func UpdatePlayer(token string, player structs.PlayerInfo) {
	db := OpenSQLite()
	defer db.Close()

	sqlStmt := `
		UPDATE wan
			SET name_adjective = ?,
				name_noun = ?,
				picture = ?
			WHERE token = ?;`
	_, err := db.Exec(sqlStmt, player.Name.Adjective, player.Name.Noun, player.Picture, token)
	
	if err != nil {
		log.Fatal(err)
	}
}

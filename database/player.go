package database

import (
	"database/sql"
	"edu/letu/wan/structs"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ClearPlayerTable() {
	db := OpenSQLite()
	defer db.Close()

	_, err := db.Exec(`DELETE FROM wan;`)
	if err != nil {
		log.Fatal(err)
	}
}

func AuthorizationExists(token, uuid string) bool {
	db := OpenSQLite()
	defer db.Close()

	sqlStmt := `
		SELECT id FROM wan 
			WHERE token = ?
			AND uuid = ?;`
	err := db.QueryRow(sqlStmt, token, uuid).Scan()

	return err != sql.ErrNoRows
}

func playerExists(db *sql.DB, token, uuid, id string) bool {
	sqlStmt := `
		SELECT id FROM wan 
			WHERE token = ?
			OR uuid = ?
			OR id = ?;`
	err := db.QueryRow(sqlStmt, token, uuid, id).Scan()

	return err != sql.ErrNoRows
}

func AddPlayer(token, uuid string, player structs.Player) bool {
	db := OpenSQLite()
	defer db.Close()

	RemovePlayerByUUID(uuid)

	if playerExists(db, token, uuid, player.ID) {
		return false
	}

	sqlStmt := `
		INSERT INTO wan (
			expires,
			token,
			uuid,
			id,
			name_adjective,
			name_noun,
			picture
		) VALUES (
			datetime('now', '1 day'),
			?, ?, ?, ?, ?, ?
		);`
	_, err := db.Exec(sqlStmt, token, uuid, player.ID, player.Name.Adjective, player.Name.Noun, player.Picture)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func rowToPlayer(row *sql.Row) *structs.Player {
	//variables to store player info
	var id string
	var name_adjective, name_noun, picture int

	//query sql
	err := row.Scan(&id, &name_adjective, &name_noun, &picture)
	if err == sql.ErrNoRows {
		return nil
	}

	//return player info
	return &structs.Player{
		ID: id,
		Name: structs.PlayerName{
			Adjective: name_adjective,
			Noun: name_noun,
		},
		Picture: picture,
	}
}

func GetPlayerByID(id string) *structs.Player {
	db := OpenSQLite()
	defer db.Close()

	sqlStmt := `
		SELECT id, name_adjective, name_noun, picture
			FROM wan
			WHERE id = ?;`

	return rowToPlayer(db.QueryRow(sqlStmt, id))
}

func GetPlayerByToken(token, uuid string) *structs.Player {
	db := OpenSQLite()
	defer db.Close()

	sqlStmt := `
		SELECT id, name_adjective, name_noun, picture
			FROM wan
			WHERE token = ?
			AND uuid = ?;`

	return rowToPlayer(db.QueryRow(sqlStmt, token, uuid))
}

func UpdatePlayer(token string, player *structs.Player) {
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

func RemovePlayerByUUID(uuid string) {
	db := OpenSQLite()
	defer db.Close()

	sqlStmt := `
		DELETE FROM wan
			WHERE uuid = ?;`
	_, err := db.Exec(sqlStmt, uuid)
	
	if err != nil {
		log.Fatal(err)
	}
}

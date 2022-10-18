package database

import (
	"database/sql"
	"edu/letu/wan/structs"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeTable() {
	
	//delete database file in case it exists
	os.Remove("./wan.db")

	db, err := sql.Open("sqlite3", "./wan.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create table, then empty in case it already exists
	sqlStmt := `
	CREATE TABLE wan (
		id TEXT NOT NULL,
		token TEXT NOT NULL,
		name_adjective INTEGER NOT NULL,
		name_noun INTEGER NOT NULL,
		picture INTEGER NOT NULL,
		PRIMARY KEY (id, token)
	);
	DELETE FROM wan;
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func SetPlayer(id string, token string, nameAdjective int, nameNoun, picture int) {
	//load db
	db, err := sql.Open("sqlite3", "./wan.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//sql statement
	sqlStmt := `INSERT INTO wan (id, token, name_adjective, name_noun, picture) VALUES (?, ?, ?, ?, ?);`

	//prepare sql query
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		fmt.Println("a")
		log.Fatal(err)
	}
	defer stmt.Close()

	//execute sql query
	var name string
	_, err = stmt.Exec(id, token, nameAdjective, nameNoun, picture)
	if err != nil {
		fmt.Println("b")
		log.Fatal(err)
	}
	fmt.Println(name)
}

func getPlayerBase(query string, value string) *structs.PlayerInfo {
	//load db
	db, err := sql.Open("sqlite3", "./wan.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//sql statement
	sqlStmt := fmt.Sprintf("SELECT id, name_adjective, name_noun, picture FROM wan WHERE %s = ?;", query)

	//prepare sql query
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//execute sql query
	var id string
	var name_adjective, name_noun, picture int
	err = stmt.QueryRow(value).Scan(&id, &name_adjective, &name_noun, &picture)
	if err != nil {
		log.Fatal(err)
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

func GetPlayer(id string) *structs.PlayerInfo {
	return getPlayerBase("id", id)
}

func GetPlayerByToken(token string) *structs.PlayerInfo {
	return getPlayerBase("token", token)
}

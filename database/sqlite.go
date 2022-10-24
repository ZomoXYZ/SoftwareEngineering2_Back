package database

import (
	"database/sql"
	"log"
)

func OpenSQLite() *sql.DB {
	db, err := sql.Open("sqlite3", "./wan.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	//create table if it doesn't exist
	err = db.QueryRow(`
		SELECT name FROM sqlite_master
			WHERE type='table'
			AND name='wan';`).Scan()
	if err == sql.ErrNoRows {
		_, err = db.Exec(`
			CREATE TABLE wan (
				uuid TEXT NOT NULL,
				id TEXT NOT NULL,
				token TEXT NOT NULL,
				expires TEXT NOT NULL,
				name_adjective INTEGER NOT NULL,
				name_noun INTEGER NOT NULL,
				picture INTEGER NOT NULL,
				PRIMARY KEY (uuid, id, token)
			);
		`)
		if err != nil {
			log.Fatal(err)
		}
	}

	//remove expired tokens
	_, err = db.Exec(`
		DELETE FROM wan
			WHERE expires < datetime('now');`)
	if err != nil {
		log.Fatal(err)
	}

	return db;
}

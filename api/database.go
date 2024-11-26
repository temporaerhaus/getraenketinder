package api

import (
	"database/sql"
	"log"
)

func InitializeDatabase(db *sql.DB){
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS getraenke (
		uuid STRING NOT NULL PRIMARY KEY, 
		name STRING, 
		presentation STRING, 
		likes INTEGER, 
		dislikes INTEGER, 
		superlikes INTEGER
	);
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER NOT NULL PRIMARY KEY,
		getraenk STRING,
		url STRING,
		hash STRING,
		height INTEGER,
		width INTEGER,
		background_hash STRING,
		mime_type STRING
	);`
	
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
}

func CreateNewGetraenk(db *sql.DB, getraenk *Getraenk) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO getraenke(uuid, name, presentation, likes, dislikes, superlikes) values (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(getraenk.UUID, getraenk.Name, getraenk.Presentation, getraenk.Likes, getraenk.Dislikes, getraenk.Superlikes)
	if err != nil {
		return err
	}
	stmt, err = tx.Prepare("INSERT INTO images(getraenk, url, hash, height, width, background_hash, mime_type) values (?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, image := range getraenk.Images{
		_, err := stmt.Exec(getraenk.UUID, image.Url, image.Hash, image.Height, image.Width, image.BackgroundHash, image.MimeType)
		if err != nil {
			return err
		}
	} 
	if err = tx.Commit(); err != nil{
		return err
	}
	return nil
}
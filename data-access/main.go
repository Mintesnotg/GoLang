package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	id     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "Minte@1997##"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "new_schema"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A single Album found: %v\n", alb)

	albID, err := addAlbum(Album{

		Title:  "Hagere",
		Artist: "Dawit Tsege",
		Price:  4.99,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id of added album : %v\n", albID)
}

func albumsByArtist(name string) ([]Album, error) {

	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist =?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {

		var alb Album

		if err := rows.Scan(&alb.id, &alb.Title, &alb.Artist, &alb.Price); err != nil {

			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil

}

func albumByID(id int64) (Album, error) {

	var alb Album

	row := db.QueryRow("SELECT * FROM album where id =?", id)

	if err := row.Scan(&alb.id, &alb.Title, &alb.Artist, &alb.Price); err != nil {

		if err == sql.ErrNoRows {

			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}

		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return alb, nil

}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// mustGetenv returns the env value for key, exiting with a clear hint if missing.

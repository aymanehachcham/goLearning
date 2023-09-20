package main

import (
	"fmt"
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}

func main() {

	config := mysql.Config{
		User: "root",
		Passwd: "",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",
		AllowNativePasswords: true,	
	}

	// database handle:
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {

		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	
	albums, err := albumsByArtist("John Coltrane")
	albumId, err := insertAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
    	Artist: "Betty Carter",
    	Price:  49.99,
	})
	fmt.Println(albumId)
	fmt.Println(albums)

}

func albumsByArtist(name string) ([]Album, error) {

	var albums []Album

	records, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer records.Close()

	for records.Next() {
		var album Album
		if err := records.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
		albums = append(albums, album)
	}

	if err := records.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func insertAlbum(album Album) (int64, error) {
	insertedAlb, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, err
	}

	idAlbum, err := insertedAlb.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insertAlbum %v", err)
	}
	return idAlbum, nil
}

func deletAlbum(id int64) error {
	_, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleteAlbum %v", err)
	}
	return nil
}
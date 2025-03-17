package database

import "database/sql"

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAll() ([]Album, error) {
	rows, err := DB.Query("SELECT id, title, artist, price FROM fakedb.album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func PostAlbum(newAlbum Album) (status int, albumData *sql.Row) {
	_, err := verifyConnectionState()
	if err != nil {
		panic(err)
	}

	result, err := DB.Exec("INSERT INTO fakedb.album (title, artist, price) VALUES (?, ?, ?)",
		newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		return 500, nil
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 500, nil
	}

	//TODO: fix this, it return a blank object
	var qResult = DB.QueryRow("SELECT * FROM fakedb.album WHERE id = ? LIMIT 1", lastInsertID)

	return 200, qResult
}

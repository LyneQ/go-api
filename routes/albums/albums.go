package albums

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"main/database"
	"net/http"
)

/*
===================================
	Start Declaration all routes
===================================
*/

// GetAlbums send the list of all albums
func GetAlbums(c *gin.Context) {
	var albums, err = database.GetAll()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, albums)
}

// PostAlbums inset the new albums to the array and return success and the created element
func PostAlbums(c *gin.Context) {
	var newAlbum database.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	var result *sql.Row
	var status int
	status, result = database.PostAlbum(newAlbum)
	if status != 200 {
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

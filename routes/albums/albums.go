package albums

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// [GO] `type` define a type for variable and `struct` tell the type is structured
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

/*
[GO] `var` define a variable, `[]album` says the variable should contain only an array with some album type
*/
var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "Jon Doe", Price: 29.99},
	{ID: 2, Title: "Kind of Blue", Artist: "Miles Davis", Price: 39.99},
	{ID: 3, Title: "Abbey Road", Artist: "The Beatles", Price: 19.99},
	{ID: 4, Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 24.99},
	{ID: 5, Title: "Back in Black", Artist: "AC/DC", Price: 15.99},
}

/*
===================================
	Start Declaration all routes
===================================
*/

// GetAlbums send the list of all albums
func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// PostAlbums inset the new albums to the array and return success and the created element
func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

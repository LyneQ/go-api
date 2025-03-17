package main

// [GO] insert in parentheses all required modules (like npm library) and install it with `go get .`
import (
	"github.com/gin-gonic/gin"
	"main/database"
	"main/routes/albums"
	"main/routes/hello"
)

func main() {
	// Creates a new instance of the Gin router with default middleware (logger and recovery)
	router := gin.Default()
	// Defines a GET HTTP route `/albums` and associates it with the handler function `getAlbums`
	router.GET("/albums", albums.GetAlbums)
	// Define a POST route `/albums` and associates it with the handler function `postAlbums`
	router.POST("/albums", albums.PostAlbums)

	router.GET("/hello", hello.SayHello)
	//Starts the HTTP server on the specified address and port.
	//err := router.Run("localhost:8080")
	//if err != nil {
	//	return
	//}

	sql.Main()

}

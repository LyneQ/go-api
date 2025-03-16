package hello

import "github.com/gin-gonic/gin"

type message struct {
	Message string `json:"message"`
}

func SayHello(c *gin.Context) {
	c.JSON(200, message{Message: "Hello there !"})
}

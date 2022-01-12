package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func main() {
	fmt.Println("Sample Web Application in Go")

	router := gin.Default()
	router.GET("/hello/:name", HelloHandler)

	router.Run(":5050")

}

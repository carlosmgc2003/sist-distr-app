package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		if len(name) == 0 {
			name = "no especificado"
		}
		c.JSON(http.StatusOK, gin.H{"nombre": name})
	})

	router.Run(":" + port)
}

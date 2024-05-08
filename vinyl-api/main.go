package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Heart Break", Artist: "DUKI", Price: 111},
	{ID: "2", Title: "Un Verano Sin Ti", Artist: "Bad Bunny", Price: 99.9},
	{ID: "3", Title: "Neo", Artist: "Sinahi", Price: 1111},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message ": "album no encontrado"})
}

// Terminar funcion de eliminar y crear función para editar
func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message ": "album no encontrado"})
}

func main() {
	fmt.Println("Bienvenido a Vinyl-Api")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

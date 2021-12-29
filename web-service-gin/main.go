package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	log.SetPrefix("Gin API: ")

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")

}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(context *gin.Context) {
	albumId := context.Param("id")

	for _, album := range albums {
		if album.ID == albumId {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

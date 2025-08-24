package main

import (
	"fmt"
	"net/http"
	"web-service-gin/database"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	database.ConnectDB()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", addAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	var albums []models.Album
	database.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Bind error:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albumID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var album models.Album
	if err := database.DB.Find(&album, "id = ?", albumID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)

}

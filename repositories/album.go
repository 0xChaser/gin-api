package repositories

import (
	"fmt"
	"net/http"
	"web-service-gin/database"
	"web-service-gin/exceptions"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	database.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Bind error:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albumID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": exceptions.InvalidUUID})
		return
	}

	var album models.Album
	if err := database.DB.First(&album, "id = ?", albumID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": exceptions.AlbumNotFound})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	albumID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": exceptions.InvalidUUID})
		return
	}

	var existingAlbum models.Album
	if err := database.DB.First(&existingAlbum, "id = ?", albumID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": exceptions.AlbumNotFound})
		return
	}

	var updateData models.Album
	if err := c.BindJSON(&updateData); err != nil {
		fmt.Println("Bind error:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData.ID = existingAlbum.ID

	if err := database.DB.Save(&updateData).Error; err != nil {
		fmt.Println("Update error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}

	c.IndentedJSON(http.StatusOK, updateData)
}

func PatchAlbumByID(c *gin.Context) {
	id := c.Param("id")

	albumID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": exceptions.InvalidUUID})
		return
	}

	var existingAlbum models.Album
	if err := database.DB.First(&existingAlbum, "id = ?", albumID).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": exceptions.AlbumNotFound})
		return
	}

	var patchData map[string]interface{}
	if err := c.BindJSON(&patchData); err != nil {
		fmt.Println("Bind error:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delete(patchData, "id")

	if err := database.DB.Model(&existingAlbum).Updates(patchData).Error; err != nil {
		fmt.Println("Patch error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}

	c.IndentedJSON(http.StatusOK, existingAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	albumID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": exceptions.InvalidUUID})
		return
	}

	result := database.DB.Delete(&models.Album{}, "id = ?", albumID)

	if result.Error != nil {
		fmt.Println("Delete error", result.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete album"})
		return
	}

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": exceptions.AlbumNotFound})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted succesfully"})
}

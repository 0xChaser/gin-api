package main

import (
	"web-service-gin/database"
	"web-service-gin/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()
	router.GET("/albums", repositories.GetAlbums)
	router.GET("/albums/:id", repositories.GetAlbumByID)
	router.POST("/albums", repositories.AddAlbums)
	router.PATCH("/albums/:id", repositories.PatchAlbumByID)
	router.PUT("/albums/:id", repositories.UpdateAlbumByID)
	router.DELETE("/albums/:id", repositories.DeleteAlbumByID)

	router.Run("localhost:8080")
}

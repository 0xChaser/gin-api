package main

import (
	"web-service-gin/database"
	"web-service-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()

	routers.AlbumRoutes(router)

	router.Run("localhost:8080")
}

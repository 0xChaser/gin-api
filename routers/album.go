package routers

import (
	"web-service-gin/repositories"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.Engine) {
	albums := router.Group("/albums")
	{
		albums.GET("", repositories.GetAlbums)
		albums.GET("/:id", repositories.GetAlbumByID)
		albums.POST("", repositories.AddAlbums)
		albums.PATCH("/:id", repositories.PatchAlbumByID)
		albums.PUT("/:id", repositories.UpdateAlbumByID)
		albums.DELETE("/:id", repositories.DeleteAlbumByID)
	}
}

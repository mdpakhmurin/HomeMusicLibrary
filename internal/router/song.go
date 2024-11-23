package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/controller"
)

type SongRouter struct{}

func (r *SongRouter) Route(route *gin.Engine) {
	Controller := controller.SongController{}
	// route.GET("/song", Controller.GetSongs)
	route.POST("/song", Controller.SongCreate)
	route.PUT("/song", Controller.SongUpdate)
	route.DELETE("/song", Controller.DeleteSong)
	route.GET("/song/search", Controller.SearchSong)
	route.GET("/song/info", Controller.GetSong)
	route.GET("/song/verses", Controller.SongVersesGet)
}

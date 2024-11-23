package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/controller/song"
)

type SongRouter struct{}

func (r *SongRouter) Route(route *gin.Engine) {
	Controller := song.SongController{}

	route.POST("/song", Controller.SongCreate)
	route.PUT("/song", Controller.SongUpdate)
	route.DELETE("/song", Controller.SongDelete)
	route.GET("/song/search", Controller.SongSearch)
	route.GET("/song/info", Controller.SongGet)
	route.GET("/song/verses", Controller.SongVersesGet)
}

package app

import (
	// "github.com/gin-gonic/gin"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mdpakhmurin/HomeMusicLibrary/cmd/docs"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	ipAddr string
	port   string
}

func NewApp(ipAddr string, port string) *App {
	return &App{
		ipAddr: ipAddr,
		port:   port,
	}
}

func (app *App) StartServer() (err error) {
	ginRoutEn := gin.Default()

	r := router.RouteLoader{}
	for _, router := range r.GetRouters() {
		router.Route(ginRoutEn)
	}

	ginRoutEn.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = ginRoutEn.Run(fmt.Sprintf("%s:%s", app.ipAddr, app.port))

	return err
}

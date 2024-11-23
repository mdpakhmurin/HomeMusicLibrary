package router

import "github.com/gin-gonic/gin"

type RouterInterface interface {
	Route(*gin.Engine)
}

type RouteLoader struct{}

func (loader RouteLoader) GetRouters() []RouterInterface {
	return []RouterInterface{
		new(SongRouter),
	}
}

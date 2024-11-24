package router

import "github.com/gin-gonic/gin"

// Интерфейс роутера
type RouterInterface interface {
	// Регистрации путей роутером
	Route(*gin.Engine)
}

// Загрузчик роутеров
type RouteLoader struct{}

// Получение роутеров
func (loader RouteLoader) GetRouters() []RouterInterface {
	return []RouterInterface{
		new(SongRouter),
	}
}

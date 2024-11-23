package app

import (
	// "github.com/gin-gonic/gin"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mdpakhmurin/HomeMusicLibrary/cmd/docs"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/router"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Экземпляр приложения
type App struct {
	ipAddr       string // IP аддрес API
	port         string // Порт API
	dbConnString string // Строка подключения к БД
}

// Создает новый экземпляр приложения
func NewApp(ipAddr string, port string, dbConnString string) *App {
	return &App{
		ipAddr:       ipAddr,
		port:         port,
		dbConnString: dbConnString,
	}
}

// Запуск приложения
func (app *App) StartApp() (err error) {
	err = repository.InitRepository(app.dbConnString)
	if err != nil {
		return errors.Wrap(err, "ошибка иникицализации репозиториев")
	}

	// Инициализация сервисов
	err = service.InitServices()
	if err != nil {
		return errors.Wrap(err, "ошибка иникицализации сервисов")
	}

	// Инициализация API
	ginEn, err := app.initApi()
	if err != nil {
		return errors.Wrap(err, "ошибка иникицализации api")
	}

	log.Info("Приложение успешно инициализировано. Запуск")
	err = ginEn.Run(fmt.Sprintf("%s:%s", app.ipAddr, app.port))
	if err != nil {
		return errors.Wrap(err, "ошибка запуска API")
	}

	return nil
}

// Инициализирует API
func (app *App) initApi() (ginEn *gin.Engine, err error) {
	ginRout := gin.Default()

	r := router.RouteLoader{}
	for _, router := range r.GetRouters() {
		router.Route(ginRout)
	}

	ginRout.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return ginRout, nil
}

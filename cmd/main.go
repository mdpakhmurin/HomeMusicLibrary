package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	_ "github.com/mdpakhmurin/HomeMusicLibrary/cmd/docs"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/app"
)

// @title Home Music Library API
// @version 1.0
// @description This is a simple API for managing home music library.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка чтение .env файла")
	}

	app := app.NewApp(os.Getenv("IP_ADDR"), os.Getenv("PORT"))
	err = app.StartServer()
	if err != nil {
		log.Fatalf("Работа сервера завершилась с ошибкой: %v", err)
	}
}

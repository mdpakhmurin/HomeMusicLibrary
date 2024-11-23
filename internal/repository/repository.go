package repository

import (
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/implementations"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/interfaces"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	SongRepository interfaces.SongRepository
)

// Инициализация репозиториев
func InitRepository(dbConnString string) (err error) {
	db, err := connectDB(dbConnString)
	if err != nil {
		return errors.Wrap(err, "ошибка иницализации БД")
	}

	err = migrateDB(db)
	if err != nil {
		return errors.Wrap(err, "ошибка иницализации БД")
	}

	err = initRepository(db)
	if err != nil {
		return errors.Wrap(err, "ошибка иницализации БД")
	}

	return nil
}

// Инициализация реализаций репозиториец
func initRepository(db *gorm.DB) (err error) {
	SongRepository = implementations.NewSongRepository(db)

	return nil
}

// Подключение к БД
func connectDB(connStr string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "ошибка подключения к БД")
	}

	return db, nil
}

// Миграция БД
func migrateDB(db *gorm.DB) (err error) {
	for _, model := range model.GetModels() {
		err = db.Debug().AutoMigrate(model.Model)
		if err != nil {
			errors.Wrap(err, "ошибка миграции БД")
		}
	}

	return nil
}

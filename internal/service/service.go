package service

import (
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/implementations"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/interfaces"
)

var (
	SongService interfaces.SongService
)

// Инициалиазция реализаций сервисов
func InitServices() (err error) {
	SongService = implementations.NewSongService()

	return nil
}

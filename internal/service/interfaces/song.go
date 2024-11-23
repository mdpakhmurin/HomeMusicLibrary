package interfaces

import "github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"

// Интерфейс сервиса песен
type SongService interface {
	// Создание песни
	// Возвращает id созданной песни
	Create(song *dto.SongCreateDtoIn) (id int, err error)
}

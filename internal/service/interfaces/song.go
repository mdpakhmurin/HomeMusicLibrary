package interfaces

import "github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"

// Интерфейс сервиса песен
type SongService interface {
	// Создание песни.
	// Возвращает id созданной песни.d
	Create(song *dto.SongCreateDtoIn) (id int, err error)

	// Получени песни по названию (точное совпадение).
	GetInfoByName(songName string, groupName string) (song *dto.SongInfoDtoOut, err error)

	// Удаление песни.
	// Возвращает id удаленной песни.
	DeleteByName(songName string, groupName string) (id int, err error)
}

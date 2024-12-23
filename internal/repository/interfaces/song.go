package interfaces

import (
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/dao"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"
)

// Интерфейс репозитория песен.
type SongRepository interface {
	// Создание песни.
	// Возвращает id созданной песни.
	Create(song *model.Song) (id int, err error)

	// Удаление песни.
	Delete(song *model.Song) (err error)

	// Обновление песни.
	Update(song *model.Song) (err error)

	// Получение песни по имение.
	GetByName(songName string, groupName string) (song *model.Song, err error)

	// Поиск песен (с пагинацией).
	// Возваращет список песен с сортировкой по названию.
	Search(searchParams *dao.SongSearchParams) (songs []model.SongShort, err error)
}

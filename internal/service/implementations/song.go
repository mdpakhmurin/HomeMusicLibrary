package implementations

import (
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"
	"github.com/pkg/errors"
)

// Сервис песен
type SongService struct {
}

// Создание нового экземпляра сервиса песен
func NewSongService() *SongService {
	return &SongService{}
}

// Создание песни
// Возвращает id созданной песни
func (service SongService) Create(song *dto.SongCreateDtoIn) (id int, err error) {
	id, err = repository.SongRepository.Create(&model.Song{
		Group:       song.Group,
		Name:        song.Name,
		Text:        song.Text,
		Link:        song.Link,
		ReleaseDate: song.ReleaseDate,
	})

	if err != nil {
		return 0, errors.Wrapf(err, "ошибка в сервисе создания песни %v", song)
	}

	return id, nil
}

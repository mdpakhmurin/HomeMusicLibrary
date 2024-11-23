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

// Создание песни.
// Возвращает id созданной песни
func (service *SongService) Create(song *dto.SongCreateDtoIn) (id int, err error) {
	id, err = repository.SongRepository.Create(&model.Song{
		Group:       song.Group,
		Name:        song.Name,
		Text:        song.Text,
		Link:        song.Link,
		ReleaseDate: song.ReleaseDate,
	})

	if err != nil {
		return 0, errors.Wrapf(err, "ошибка в сервисе, при создании песни %v", song)
	}

	return id, nil
}

// Удаление песни. Возвращает id удаленной песни
func (service *SongService) DeleteByName(songName string, groupName string) (id int, err error) {
	songModel, err := repository.SongRepository.GetByName(songName, groupName)
	if err != nil {
		return 0, errors.Wrapf(err, "ошибка в сервисе, при удалении песни %v", songName)
	}

	err = repository.SongRepository.Delete(songModel)
	if err != nil {
		return 0, errors.Wrapf(err, "ошибка в сервисе, при удалении песни %v", songName)
	}

	return songModel.Id, nil
}

// Получение песни по названию (точное совпадение)
func (service *SongService) GetInfoByName(songName string, groupName string) (song *dto.SongInfoDtoOut, err error) {
	songModel, err := repository.SongRepository.GetByName(songName, groupName)
	if err != nil {
		return nil, errors.Wrapf(err, "ошибка в сервисе, при удалении песни %v", song)
	}

	return &dto.SongInfoDtoOut{
		Id:          songModel.Id,
		Text:        songModel.Text,
		Link:        songModel.Link,
		ReleaseDate: songModel.ReleaseDate,
	}, nil
}

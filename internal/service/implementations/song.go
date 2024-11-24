package implementations

import (
	"strings"

	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"
	"github.com/mdpakhmurin/HomeMusicLibrary/pkg/mathutils.go"
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

// Обновление песни с таким же названием и группой.
// Возвращает id обновлённой записи.
func (service *SongService) Update(song *dto.SongUpdateDtoIn) (id int, err error) {
	songModel, err := repository.SongRepository.GetByName(song.Name, song.Group)
	if err != nil {
		return 0, errors.Wrapf(err, "ошибка в сервисе, при обновлении песни %v", song)
	}

	songModel.Text = song.Text
	songModel.Link = song.Link
	songModel.ReleaseDate = song.ReleaseDate

	err = repository.SongRepository.Update(songModel)
	if err != nil {
		return 0, errors.Wrapf(err, "ошибка в сервисе, при обновлении песни %v", song)
	}

	return songModel.Id, nil
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
		return nil, errors.Wrapf(err, "ошибка в сервисе, при получении песни %v", song)
	}

	return &dto.SongInfoDtoOut{
		Id:          songModel.Id,
		Text:        songModel.Text,
		Link:        songModel.Link,
		ReleaseDate: songModel.ReleaseDate,
	}, nil
}

// Получение куплетов песни с пагинацией
func (service *SongService) SongVersesGet(song *dto.SongGetVersesDtoIn) ([]string, error) {
	songModel, err := repository.SongRepository.GetByName(song.Name, song.Group)
	if err != nil {
		return nil, errors.Wrapf(err, "ошибка в сервисе, при получении текста песни %v", song)
	}

	verses := strings.Split(songModel.Text, "\n\n")

	page := song.Page
	pageSize := song.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 0
	}

	firstVerseId := (page - 1) * pageSize
	firstVerseId = mathutils.BoundNumber(firstVerseId, 0, len(verses))

	endVerseId := (page * song.PageSize)
	endVerseId = mathutils.BoundNumber(endVerseId, 0, len(verses))

	return verses[firstVerseId:endVerseId], nil
}

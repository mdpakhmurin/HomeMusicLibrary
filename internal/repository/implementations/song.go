package implementations

import (
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/dao"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Репозиторий песен.
type SongRepository struct {
	db *gorm.DB
}

// Создание объекта репозитория песен.
func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

// Создание новой песни.
func (s *SongRepository) Create(song *model.Song) (id int, err error) {
	if err := s.db.Create(&song).Error; err != nil {
		return 0, errors.Wrapf(err, "ошибка в репозитории песни, при создании сущности песни %v", song)
	}

	log.Infof("Сущность песни успешно создана (%#v)", song)

	return song.Id, nil
}

// Получение песни по названию песни и группы (точное совпадение).
func (s *SongRepository) GetByName(songName string, groupName string) (song *model.Song, err error) {
	song = &model.Song{
		Name:  songName,
		Group: groupName,
	}
	result := s.db.Where(song).First(song)

	err = result.Error
	if err != nil {
		return nil, errors.Wrapf(err, "ошибка в репозитории песни, при получение сущности по названию %v", song)
	}

	return song, err
}

// Обновление песни.
func (s *SongRepository) Update(song *model.Song) (err error) {
	result := s.db.Updates(song)
	err = result.Error
	if err != nil {
		return errors.Wrapf(err, "ошибка в репозитории песни, при обновлении сущности  %v", song)
	}

	log.Infof("Сущность песни успешно обновлена (%#v)", song)

	return nil
}

// Удаление песни.
func (s *SongRepository) Delete(song *model.Song) (err error) {
	result := s.db.Where(song).Delete(song)
	err = result.Error
	if err != nil {
		return errors.Wrapf(err, "ошибка в репозитории песни, при удалении сущности  %v", song)
	}

	log.Infof("Сущность песни успешно удалена (%#v)", song)

	return nil
}

// Поиск песен по параметрам (с пагинацией).
// Возвращает список песен с сортировкой по названию.
func (s *SongRepository) Search(searchParams *dao.SongSearchParams) (songs []model.SongShort, err error) {
	// Подгатовка параметров пагинации
	page := searchParams.Page
	pageSize := searchParams.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 0
	}

	// Составление запроса фильтрации
	query := s.db.
		Where("group_name like ?", "%"+searchParams.Group+"%").
		Where("song_name like ?", "%"+searchParams.Name+"%").
		Where("song_text like ?", "%"+searchParams.Text+"%").
		Where("song_link like ?", "%"+searchParams.Link+"%")
	if searchParams.ReleaseStartDate != nil {
		query = query.Where("release_date >= ?", searchParams.ReleaseStartDate)
	}
	if searchParams.ReleaseEndDate != nil {
		query = query.Where("release_date <= ?", searchParams.ReleaseEndDate)
	}

	// Выполнение запроса с пагинацией
	result := query.
		Order("song_name").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&songs)

	err = result.Error
	if err != nil {
		return nil, errors.Wrapf(err, "ошибка в репозитории песни, при поиске сущностей")
	}
	return songs, nil
}

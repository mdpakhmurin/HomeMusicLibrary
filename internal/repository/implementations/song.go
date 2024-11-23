package implementations

import (
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Репозиторий песен
type SongRepository struct {
	db *gorm.DB
}

// Создание объекта репозитория песен
func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

// Создание новой песни
func (s SongRepository) Create(song *model.Song) (id int, err error) {
	if err := s.db.Create(&song).Error; err != nil {
		return 0, errors.Wrapf(err, "ошибка в репозитории создания сущности песни %v", song)
	}

	log.Infof("Сущность песни успешно создана (%v)", song)

	return song.Id, nil
}

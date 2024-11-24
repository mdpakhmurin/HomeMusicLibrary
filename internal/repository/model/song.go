package model

import (
	"time"
)

// Модель песни.
type Song struct {
	Id          int        `gorm:"primaryKey;autoIncrement;column:id"`                     // ID Песни.
	Group       string     `gorm:"column:group_name;uniqueIndex:idx_group_name_song_name"` // название группы.
	Name        string     `gorm:"column:song_name;uniqueIndex:idx_group_name_song_name"`  // Название песни.
	Text        string     `gorm:"column:song_text"`                                       // Текст песни.
	Link        string     `gorm:"column:song_link"`                                       // Ссылка на песню.
	ReleaseDate *time.Time `gorm:"column:release_date"`                                    //Дата создания песни.
}

// Имя таблицы в БД.
func (Song) TableName() string {
	return "song"
}

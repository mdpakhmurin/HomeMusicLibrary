package model

import (
	"time"
)

// Модель песни
type Song struct {
	Id          int        `gorm:"primaryKey;autoIncrement;column:id"`
	Group       string     `gorm:"column:group_name;uniqueIndex:idx_group_name_song_name"`
	Name        string     `gorm:"column:song_name;uniqueIndex:idx_group_name_song_name"`
	Text        string     `gorm:"column:song_text"`
	Link        string     `gorm:"column:song_link"`
	ReleaseDate *time.Time `gorm:"column:release_date"`
}

// Имя таблицы в БД
func (Song) TableName() string {
	return "song"
}

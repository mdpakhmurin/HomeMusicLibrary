package model

import "time"

// Сокращенная модель песни.
type SongShort struct {
	Id          int        `gorm:"column:id"`           // ID песни.
	Group       string     `gorm:"column:group_name"`   // Название группы.
	Name        string     `gorm:"column:song_name"`    // Название песни.
	Link        string     `gorm:"column:song_link"`    // Ссылка на песню.
	ReleaseDate *time.Time `gorm:"column:release_date"` // Дата создания песни.
}

// Имя таблицы.
func (SongShort) TableName() string {
	return "song"
}

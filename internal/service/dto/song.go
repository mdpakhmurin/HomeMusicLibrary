package dto

import "time"

// DTO создания песни
type SongCreateDtoIn struct {
	Group       string
	Name        string
	Text        string
	Link        string
	ReleaseDate time.Time
}

// DTO обновления песни
type SongUpdateDtoIn struct {
	Group       string
	Name        string
	Text        string
	Link        string
	ReleaseDate time.Time
}

// DTO поиска песен
type SongSearchDtoIn struct {
	Group       string
	Name        string
	Text        string
	Link        string
	ReleaseDate string
	Page        int
	PageSize    int
}

// DTO-ответ поиска песни
type SongSearchDtoOut struct {
	Id          int
	Group       string
	Name        string
	ReleaseDate time.Time
	Link        string
}

// DTO информации песни
type SongInfoDtoIn struct {
	Group string
	Name  string
}

// DTO-ответ информации о песни
type SongInfoDtoOut struct {
	Id          int
	Text        string
	Link        string
	ReleaseDate time.Time
}

// DTO получения текста песни с пагинацией по куплетам
type SongGetVersesDtoIn struct {
	Group    string
	Name     string
	Page     int
	PageSize int
}

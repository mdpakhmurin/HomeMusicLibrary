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
	Group       string `json:"group"`
	Name        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}

// DTO удаления песни
type SongRemoveDtoIn struct {
	Group string `form:"group"`
	Name  string `form:"song"`
}

// DTO поиска песен
type SongSearchDtoIn struct {
	Group       string `form:"group"`
	Name        string `form:"song"`
	Text        string `form:"text"`
	Link        string `form:"link"`
	ReleaseDate string `form:"releaseDate"`
	Page        int    `form:"page"`
	PageSize    int    `form:"pageSize"`
}

// DTO-ответ поиска песни
type SongSearchDtoOut struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
}

// DTO информации песни
type SongInfoDtoIn struct {
	Group string `form:"group"`
	Name  string `form:"song"`
}

// DTO-ответ информации о песни
type SongInfoDtoOut struct {
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}

// DTO получения текста песни с пагинацией по куплетам
type SongGetVersesDtoIn struct {
	Group    string `form:"group"`
	Name     string `form:"song"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

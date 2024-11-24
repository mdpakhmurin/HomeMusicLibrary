package dto

import "time"

// DTO создания песни.
type SongCreateDtoIn struct {
	Group       string     // Название группы.
	Name        string     // Название песни.
	Text        string     // Текст песни.
	Link        string     // Ссылка на песню.
	ReleaseDate *time.Time // Дата создания песни.
}

// DTO обновления песни.
type SongUpdateDtoIn struct {
	Group       string     // Название группы.
	Name        string     // Название песни.
	Text        string     // Текст песни.
	Link        string     // Ссылка на песню.
	ReleaseDate *time.Time // Дата создания песни
}

// DTO поиска песен (с пагинацией).
type SongSearchDtoIn struct {
	Group            string     // Текст встречающийся в названии группы.
	Name             string     // Текст встречающийся в названии песни.
	Text             string     // Текст встречающийся в тексте песни.
	Link             string     // Текст встречающийся в ссылке.
	ReleaseStartDate *time.Time // Начальная дата периода поиска песни.
	ReleaseEndDate   *time.Time // Конечная дата периода поиска песни.
	Page             int        // Номер страницы.
	PageSize         int        // Размер страницы.
}

// DTO-ответ поиска песни.
type SongSearchDtoOut struct {
	Id          int        // ID песни.
	Group       string     // Название группы.
	Name        string     // Название песни.
	ReleaseDate *time.Time // Дата создания песни.
	Link        string     // Ссылка на песню.
}

// DTO информации песни.
type SongInfoDtoIn struct {
	Group string // Название группы.
	Name  string // Название песни.
}

// DTO-ответ информации о песни.
type SongInfoDtoOut struct {
	Id          int        // ID песни.
	Text        string     // Текст песни.
	Link        string     // Ссылка на песню.
	ReleaseDate *time.Time // Дата создания песни.
}

// DTO получения текста песни с пагинацией по куплетам.
type SongGetVersesDtoIn struct {
	Group    string // Название группы.
	Name     string //  Название песни.
	Page     int    // Номер страницы.
	PageSize int    // Размер страницы.
}

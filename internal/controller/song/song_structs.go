package song

// DTO создания песни
type SongCreateDtoIn struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
} //@name SongCreate

// DTO обновления песни
type SongUpdateDtoIn struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
} //@name SongUpdate

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
} //@name SongSearchFilter

// DTO-ответ поиска песни
type SongSearchDtoOut struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
} // @name SongDescription

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
} // @name SongDetail

// DTO получения текста песни с пагинацией по куплетам
type SongGetVersesDtoIn struct {
	Group    string `form:"group"`
	Name     string `form:"song"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

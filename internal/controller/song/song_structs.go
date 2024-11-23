package song

// View создания песни
type SongCreateViewIn struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
} //@name SongCreate

// View обновления песни
type SongUpdateViewIn struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
} //@name SongUpdate

// View удаления песни
type SongDeleteViewIn struct {
	Group string `form:"group"`
	Name  string `form:"song"`
}

// View поиска песен
type SongSearchViewIn struct {
	Group       string `form:"group"`
	Name        string `form:"song"`
	Text        string `form:"text"`
	Link        string `form:"link"`
	ReleaseDate string `form:"releaseDate"`
	Page        int    `form:"page"`
	PageSize    int    `form:"pageSize"`
} //@name SongSearchFilter

// View-ответ поиска песни
type SongSearchViewOut struct {
	Group       string `json:"group"`
	Name        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
} // @name SongDescription

// View информации песни
type SongInfoViewIn struct {
	Group string `form:"group"`
	Name  string `form:"song"`
}

// View-ответ информации о песни
type SongInfoViewOut struct {
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
} // @name SongDetail

// View получения текста песни с пагинацией по куплетам
type SongGetVersesViewIn struct {
	Group    string `form:"group"`
	Name     string `form:"song"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

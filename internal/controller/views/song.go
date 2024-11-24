package views

// View создания песни.
type SongCreateViewIn struct {
	Group       string `json:"group" binding:"required"` // Название группы (обязательно).
	Name        string `json:"song" binding:"required"`  // Название песни (обязательно).
	Text        string `json:"text"`                     // Текст песни.
	Link        string `json:"link"`                     // Ссылка на песню.
	ReleaseDate string `json:"releaseDate"`              // Дата создания песни.
} //@name SongCreate

// View обновления песни.
type SongUpdateViewIn struct {
	Group       string `json:"group" binding:"required"` // Название группы (обязательно).
	Name        string `json:"song" binding:"required"`  // Название песни (обязательно).
	Text        string `json:"text"`                     // Текст песни.
	Link        string `json:"link"`                     // Ссылка на песню.
	ReleaseDate string `json:"releaseDate"`              // Дата создания песни.
} //@name SongUpdate

// View удаления песни.
type SongDeleteViewIn struct {
	Group string `form:"group" binding:"required"` // Название группы (обязательно).
	Name  string `form:"song" binding:"required"`  // Название песни (обязательно).
}

// View поиска песен (с пагинацией).
type SongSearchViewIn struct {
	Group            string `form:"group"`            // Текст встречающийся в названии группы.
	Name             string `form:"song"`             // Текст встречающийся в названии песни.
	Text             string `form:"text"`             // Текст встречающийся в тексте песни.
	Link             string `form:"link"`             // Текст встречающийся в ссылке.
	ReleaseStartDate string `form:"releaseStartDate"` // Начальная дата периода поиска песни.
	ReleaseEndDate   string `form:"releaseEndDate"`   // Конечная дата периода поиска песни.
	Page             int    `form:"page"`             // Номер страницы.
	PageSize         int    `form:"pageSize"`         // Размер страницы.
} //@name SongSearchFilter

// View-ответ поиска песни.
type SongSearchViewOut struct {
	Group       string `json:"group"`       // Название группы.
	Name        string `json:"song"`        // Название песни.
	ReleaseDate string `json:"releaseDate"` // Дата создания песни.
	Link        string `json:"link"`        // Ссылка на песню.
} // @name SongDescription

// View информации песни.
type SongInfoViewIn struct {
	Group string `form:"group" binding:"required"` // Название группы (обязательно).
	Name  string `form:"song" binding:"required"`  // Название песни (обязательно).
}

// View-ответ информации о песни.
type SongInfoViewOut struct {
	Text        string `json:"text"`        // Текст песни.
	Link        string `json:"link"`        // Ссылка на песню.
	ReleaseDate string `json:"releaseDate"` // Дата релиза песни.
} // @name SongDetail

// View получения текста песни с пагинацией по куплетам.
type SongGetVersesViewIn struct {
	Group    string `form:"group" binding:"required"` // Название группы (обязательно).
	Name     string `form:"song" binding:"required"`  // Название песни (обязательно).
	Page     int    `form:"page"`                     // Номер страницы.
	PageSize int    `form:"pageSize"`                 // Размер страницы.
}

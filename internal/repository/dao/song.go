package dao

import "time"

// DAO поиска песен (с пагинацией).
type SongSearchParams struct {
	Group            string     // Текст встречающийся в названии группы.
	Name             string     // Текст встречающийся в названии песни.
	Text             string     // Текст встречающийся в тексте песни.
	Link             string     // Текст встречающийся в ссылке.
	ReleaseStartDate *time.Time // Начальная дата периода поиска песни.
	ReleaseEndDate   *time.Time // Конечная дата периода поиска песни.
	Page             int        // Номер страницы.
	PageSize         int        // Размер страницы.
}

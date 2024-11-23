package interfaces

import "github.com/mdpakhmurin/HomeMusicLibrary/internal/repository/model"

// Интерфейс репозитория песен
type SongRepository interface {
	Create(song *model.Song) (id int, err error)
}

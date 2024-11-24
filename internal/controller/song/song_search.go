package song

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/controller/views"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"
	log "github.com/sirupsen/logrus"
)

// SongSearch поиск песни.
// @Summary Поиск песен с пагинацией.
// @Description Поиск песке по заданным параметрам. Вовзращает список песен, отсортированный по названию.
// @Produce json
// @Param song query SongSearchFilter true "Данные о песне"
// @Success 200  "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /song/search [get]
func (controller *SongController) SongSearch(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := getGeneralRequestInfo(c)

	// Получение входных данных
	var input views.SongSearchViewIn
	err := getRequestData(c, &input)
	if err != nil {
		return
	}

	log.Infof("%s. (%#v)", generalLog, input)

	// Поиск песен
	songs, err := songSearch(c, &input)
	if err != nil {
		return
	}

	// Преобразование списка структур песен в формат JSON
	songsJSON, err := json.Marshal(songs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при конвертации в JSON"})
		log.Debugf("%s. Ошибка конвертации ответа в json (%#v) -> (%#v): %v", generalLog, input, songs, err)
		return
	}

	// Отправка данных JSON в ответе
	c.Data(http.StatusOK, "application/json", songsJSON)
	log.Infof("%s. Успешный ответ: %v", generalLog, string(songsJSON))
}

// Поиск песен с помощью сервсива.
func songSearch(c *gin.Context, songView *views.SongSearchViewIn) (songs []*views.SongSearchViewOut, err error) {
	// Получение дат
	startDate, err := parseDate(c, songView.ReleaseStartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := parseDate(c, songView.ReleaseEndDate)
	if err != nil {
		return nil, err
	}

	// Конвертация в DTO
	songDto := dto.SongSearchDtoIn{
		Group:            songView.Group,
		Name:             songView.Name,
		Text:             songView.Text,
		Link:             songView.Link,
		ReleaseStartDate: startDate,
		ReleaseEndDate:   endDate,
		Page:             songView.Page,
		PageSize:         songView.PageSize,
	}

	// Поиск песен
	songsDtos, err := service.SongService.SongSearch(&songDto)
	if err != nil {
		log.Debugf("%s. Ошибка поиска песен (%#v): %v.", getGeneralRequestInfo(c), songView, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка поиска"})
		return nil, err
	}

	// Конвертация DTO в структуры ответа
	for _, dto := range songsDtos {
		songs = append(songs, &views.SongSearchViewOut{
			Group:       dto.Group,
			Name:        dto.Name,
			ReleaseDate: dateToString(dto.ReleaseDate),
			Link:        dto.Link,
		})
	}

	return songs, nil
}

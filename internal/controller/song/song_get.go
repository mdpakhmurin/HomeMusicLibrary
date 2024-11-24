package song

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	log "github.com/sirupsen/logrus"
)

// @Summary Получение информации о песни
// @Description Возвращает информацию о песни на основе входных данных
// @ID SongInfo
// @Accept json
// @Produce json
// @Param input query SongInfoViewIn true "Входные данные для получения информации о песне"
// @Success 200  "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /song/info [get]
func (controller *SongController) SongInfo(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := getGeneralRequestInfo(c)

	// Получение входных данных
	var input SongInfoViewIn
	err := getRequestData(c, &input)
	if err != nil {
		return
	}

	log.Infof("%s. (%#v)", generalLog, input)

	// Получение информации о песни
	song, err := getSongInfo(c, &input)
	if err != nil {
		return
	}

	// Преобразование структуры песни в формат JSON
	songJSON, err := json.Marshal(song)
	if err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке конвертации
		log.Debugf("%s. Ошибка конвертации ответа в JSON (%#v): %v", generalLog, song, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка конвертации ответа в JSON"})
		return
	}

	// Отправка данных JSON в ответе
	c.Data(http.StatusOK, "application/json", songJSON)
	log.Infof("%s. Успешный ответ: %#v", generalLog, input)
}

// Получение  песни с помощью сервиса
func getSongInfo(c *gin.Context, songView *SongInfoViewIn) (songResponseView *SongInfoViewOut, err error) {
	generalLog := getGeneralRequestInfo(c)

	song, err := service.SongService.GetInfoByName(songView.Name, songView.Group)
	if err != nil {
		log.Debugf("%s. Ошибка получения песни (%#v): %v.", generalLog, songView, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка получения песни (%s: %s)", songView.Group, songView.Name)})
		return nil, err
	}

	return &SongInfoViewOut{
		Text:        song.Text,
		Link:        song.Link,
		ReleaseDate: song.ReleaseDate.Format("02.01.2006"),
	}, nil
}

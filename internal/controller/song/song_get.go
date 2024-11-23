package song

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SongInfo "Получение песни по группе и названию песни"
// @Summary Получение информации о песне
// @Description Получение информации о песне по заданной группе и названию
// @Produce json
// @Param song query SongInfoDtoIn true "Данные о песне"
// @Success 200 {object} SongDetail
// @Router /song/info [get]
func (controller *SongController) SongGet(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных песни
	var input SongInfoDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		return
	}

	// Создание структуры песни с данными
	song := SongInfoDtoOut{
		Text:        "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Link:        "https://2323",
		ReleaseDate: "02.08.2009",
	}

	// Преобразование структуры песни в формат JSON
	songJSON, err := json.Marshal(song)
	if err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке конвертации
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка конвертации ответа в JSON"})
		log.Debugf("%s. Ошибка конвертации ответа (%v): %v", generalLog, song, err)
		return
	}

	// Отправка данных JSON в ответе
	c.Data(http.StatusOK, "application/json", songJSON)
	log.Infof("%s. Успешный ответ: %v", generalLog, string(songJSON))
}

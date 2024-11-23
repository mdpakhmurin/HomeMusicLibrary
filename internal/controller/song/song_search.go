package song

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SongSearch поиск песни.
// @Summary Поиск песни
// @Description Поиск песни по заданным параметрам
// @Produce json
// @Param song query SongSearchFilter true "Данные о песне"
// @Success 200 {array} SongDescription
// @Router /song/search [get]
func (controller *SongController) SongSearch(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных поиска песни
	var input SongSearchViewIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		return
	}

	// Создание списка структур песен с данными
	songs := []SongSearchViewOut{
		{Group: "cool group", Name: "super song", ReleaseDate: "23.11.2024", Link: "https://"},
		{Group: "nice gr", Name: "ok song", ReleaseDate: "10.11.2024", Link: "https://new"},
	}

	// Преобразование списка структур песен в формат JSON
	songsJSON, err := json.Marshal(songs)
	if err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке конвертации
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при конвертации в JSON"})
		log.Debugf("%s. Ошибка конвертации ответа (%v): %v", generalLog, songs, err)
		return
	}

	// Отправка данных JSON в ответе
	c.Data(http.StatusOK, "application/json", songsJSON)
	log.Infof("%s. Успешный ответ: %v", generalLog, string(songsJSON))
}

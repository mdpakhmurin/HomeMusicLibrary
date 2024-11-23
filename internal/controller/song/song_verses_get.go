package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SongVersesGet получение текста песни.
// @Summary Получение текста песни
// @Description Получение куплетов песни по заданным параметрам
// @Produce json
// @Param song query SongGetVersesDtoIn true "Данные о песне"
// @Success 200 {array} string
// @Router /song/verses [get]
func (controller *SongController) SongVersesGet(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных получения куплетов песни
	var input SongGetVersesDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// Обработка ошибки и отправка ответа с кодом ошибки
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Содержание куплетов песни
	verses := []string{
		"Первый куплет текста песни",
		"Припев текста песни",
		"Второй куплет текста песни",
	}

	// Отправка ответа с куплетами песни
	c.JSON(http.StatusOK, gin.H{"verses": verses})
	log.Infof("%s. Успешный ответ: %v", generalLog, verses)
}

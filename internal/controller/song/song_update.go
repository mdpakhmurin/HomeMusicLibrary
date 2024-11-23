package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary Обновляет песню с заданным названием и группой
// @Description Добавляет новую песню
// @Accept json
// @Produce json
// @Param song body SongUpdate true "Данные о песне в формате JSON"
// @Success 200 {string} string "Сообщение о успешном добавлении песни"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Router /song/ [put]
func (controller *SongController) SongUpdate(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s", c.Request.URL.Path)

	// Инициализация переменной для входных данных обновления песни
	var input SongUpdateViewIn

	// Привязка входных данных к переменной input из JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		return
	}

	// Обновление общего лога с полученными данными
	generalLog = fmt.Sprintf("%s. %v", generalLog, input)
	log.Info(generalLog)

	// Формирование ответного сообщения об успешном обновлении песни
	answer := fmt.Sprintf("Песня (%s: %s) успешно обновлена", input.Group, input.Name)

	// Отправка ответа с сообщением об успешном обновлении
	c.JSON(http.StatusOK, gin.H{"message": answer})
	log.Infof("%s. Успешное обновление", generalLog)
}

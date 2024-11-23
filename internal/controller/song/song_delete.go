package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SongDelete удаляет песню по указанным параметрам.
// @Summary Удалить песню
// @Description Удаляет песню по указанным параметрам: group и name.
// @Produce json
// @Param song query SongRemoveDtoIn true "Данные о песне"
// @Success 200 {object} string "Сообщение об успешном удалении песни"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Router /song/ [delete]
func (controller *SongController) SongDelete(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных удаления песни
	var input SongRemoveDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Формирование строки с сообщением об успешном удалении песни
	var responseString = fmt.Sprintf("Песня (%s: %s) успешно удалена", input.Group, input.Name)

	// Отправка ответа с сообщением об успешном удалении
	c.JSON(http.StatusOK, gin.H{"message": responseString})
	log.Infof("%s. Успешное удаление: %v", generalLog, responseString)
}

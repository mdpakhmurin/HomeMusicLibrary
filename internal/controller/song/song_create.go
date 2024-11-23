package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary Добавить новую песню
// @Description Добавляет новую песню
// @Accept json
// @Produce json
// @Param song body SongCreate true "Данные о песне в формате JSON"
// @Success 200 {string} string "Сообщение о успешном добавлении песни"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Router /song [post]
func (controller *SongController) SongCreate(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s", c.Request.URL.Path)

	// Инициализация переменной для входных данных создания песни
	var input SongCreateDtoIn

	// Привязка входных данных к переменной input из JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновление общего лога с полученными данными
	generalLog = fmt.Sprintf("%s. %v", generalLog, input)
	log.Info(generalLog)

	// Формирование ответного сообщения об успешном добавлении песни
	answer := fmt.Sprintf("Песня (%s: %s) успешно добавлена", input.Group, input.Name)

	// Отправка ответа с сообщением об успешном добавлении
	c.JSON(http.StatusOK, gin.H{"message": answer})
	log.Infof("%s. Успешное добавление", generalLog)
}

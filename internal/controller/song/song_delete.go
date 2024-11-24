package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	log "github.com/sirupsen/logrus"
)

// @Summary Удаление песни
// @Description Удаляет песню по указанному идентификатору
// @ID SongDelete
// @Accept json
// @Produce json
// @Param input query SongDeleteViewIn true "Входные данные для удаления песни"
// @Success 200 "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /song/ [delete]
func (controller *SongController) SongDelete(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := getGeneralRequestInfo(c)

	// Инициализация переменной для входных данных удаления песни
	var input SongDeleteViewIn
	err := getRequestData(c, &input)
	if err != nil {
		return
	}

	log.Infof("%s. (%#v)", generalLog, input)

	// Удаление песни
	songId, err := songDelete(c, &input)
	if err != nil {
		return
	}

	// Отправка ответа с сообщением об успешном удалении
	c.JSON(http.StatusOK, gin.H{"id": songId})
	log.Infof("%s. Успешное удаление песни с id %d: %#v", generalLog, songId, input)
}

// Удаление песни с помощью сервиса
func songDelete(c *gin.Context, songView *SongDeleteViewIn) (id int, err error) {

	// Удаление песни
	id, err = service.SongService.DeleteByName(songView.Name, songView.Group)
	if err != nil {
		log.Debugf("%s. Ошибка удаления песни (%#v): %v.", getGeneralRequestInfo(c), songView, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка удаления песни (%s-%s).", songView.Name, songView.Group)})
		return
	}

	return id, nil
}

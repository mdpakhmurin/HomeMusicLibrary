package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"
	log "github.com/sirupsen/logrus"
)

// SongVersesGet получение куплетов песни.
// @Summary Получение куплетов песни
// @Description Получение куплетов песни с пагинацией
// @Produce json
// @Param song query SongGetVersesViewIn true "Данные о песне"
// @Success 200  "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /song/verses [get]
func (controller *SongController) SongVersesGet(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := getGeneralRequestInfo(c)

	// Получение входных данных
	var input SongGetVersesViewIn
	err := getRequestData(c, &input)
	if err != nil {
		return
	}

	log.Infof("%s. (%#v)", generalLog, input)

	// Получение куплетов песни
	verses, err := songVersesGet(c, &input)
	if err != nil {
		return
	}

	// Отправка ответа с куплетами песни
	c.JSON(http.StatusOK, gin.H{"verses": verses})
	log.Infof("%s. Успешный ответ: %#v", generalLog, verses)
}

// Получение куплетов песни с помощью сервиса
func songVersesGet(c *gin.Context, songView *SongGetVersesViewIn) (verses []string, err error) {
	// Конвертация в DTO
	songDto := dto.SongGetVersesDtoIn{
		Group:    songView.Group,
		Name:     songView.Name,
		Page:     songView.Page,
		PageSize: songView.PageSize,
	}

	// Получение куплетов песни
	verses, err = service.SongService.SongVersesGet(&songDto)
	if err != nil {
		log.Debugf("%s. Ошибка получение куплетов песни песни (%#v): %v.", getGeneralRequestInfo(c), songView, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка получения куплетов песни (%s: %s).", songDto.Group, songDto.Name)})
		return nil, err
	}

	return verses, nil
}

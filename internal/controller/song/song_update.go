package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"
	log "github.com/sirupsen/logrus"
)

// @Summary Обновляет песню
// @Description Обновляет песню с таким же названием и группой
// @Accept json
// @Produce json
// @Param song body SongUpdate true "Входные данные о песне"
// @Success 200  "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /song/ [put]
func (controller *SongController) SongUpdate(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := getGeneralRequestInfo(c)

	// Получение входных данных
	var input SongUpdateViewIn
	err := getBodyData(c, &input)
	if err != nil {
		return
	}

	log.Infof("%s. (%#v)", generalLog, input)

	// Обновление песни
	id, err := songUpdate(c, &input)
	if err != nil {
		return
	}

	// Отправка ответа с ID
	c.JSON(http.StatusOK, gin.H{"id": id})
	log.Infof("%s. Успешное обновление песни %#v", generalLog, input)
}

// Обновление песни с помощью сервиса
func songUpdate(c *gin.Context, songView *SongUpdateViewIn) (id int, err error) {
	// Получение даты
	releaseDate, err := parseDate(c, songView.ReleaseDate)
	if err != nil {
		return 0, err
	}

	// Конвертация в DTO
	songDto := dto.SongUpdateDtoIn{
		Group:       songView.Group,
		Name:        songView.Name,
		Text:        songView.Text,
		Link:        songView.Link,
		ReleaseDate: releaseDate,
	}

	// Обновление песни
	id, err = service.SongService.Update(&songDto)
	if err != nil {
		log.Debugf("%s. Ошибка обновления песни (%#v): %v.", getGeneralRequestInfo(c), songView, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка обновления песни (%s: %s).", songDto.Group, songDto.Name)})
		return
	}

	return id, nil
}

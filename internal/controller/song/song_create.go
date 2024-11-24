package song

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service"
	"github.com/mdpakhmurin/HomeMusicLibrary/internal/service/dto"
	log "github.com/sirupsen/logrus"
)

// @Summary Создание песни
// @Description Создает новую песню на основе предоставленных входных данных
// @ID SongCreate
// @Accept json
// @Produce json
// @Param input body SongCreateViewIn true "Входные данные для создания песни"
// @Success 200  "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /song [post]
func (controller *SongController) SongCreate(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := getGeneralRequestInfo(c)

	// Получение входных данных
	var input SongCreateViewIn
	err := getBodyData(c, &input)
	if err != nil {
		return
	}

	log.Infof("%s. (%#v)", generalLog, input)

	// Создание песни
	id, err := songCreate(c, &input)
	if err != nil {
		return
	}

	// Отправка ответа с ID
	c.JSON(http.StatusOK, gin.H{"id": id})
	log.Infof("%s. Успешное добавление песни: %#v", generalLog, input)
}

// Создание песни с помощью сервиса
func songCreate(c *gin.Context, songView *SongCreateViewIn) (id int, err error) {
	// Получение даты
	releaseDate, err := parseDate(c, songView.ReleaseDate)
	if err != nil {
		return 0, err
	}

	// Конвертация в DTO
	songDto := dto.SongCreateDtoIn{
		Group:       songView.Group,
		Name:        songView.Name,
		Text:        songView.Text,
		Link:        songView.Link,
		ReleaseDate: releaseDate,
	}

	// Создание песни
	id, err = service.SongService.Create(&songDto)
	if err != nil {
		log.Debugf("%s. Ошибка создания песни (%#v): %v.", getGeneralRequestInfo(c), songView, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка создания песни (%s: %s).", songDto.Group, songDto.Name)})
		return 0, err
	}

	return id, nil
}

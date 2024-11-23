package song

import (
	"fmt"
	"net/http"
	"time"

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
	log.Infof(generalLog)

	// Получение входных данных
	var input SongCreateViewIn
	err := getBodyData(c, &input)
	if err != nil {
		return
	}

	// Создание песни
	id, err := songCreate(c, &input)
	if err != nil {
		return
	}

	// Отправка ответа с ID
	c.JSON(http.StatusOK, gin.H{"message": id})
	log.Infof("%s. Успешное добавление песни (id %d) %s: %s", generalLog, id, input.Group, input.Name)
}

// Создание песни с помощью сервиса
func songCreate(c *gin.Context, songView *SongCreateViewIn) (id int, err error) {
	generalLog := getGeneralRequestInfo(c)

	// Получение даты
	releaseDate, err := time.Parse("02.01.2006", songView.ReleaseDate)
	if err != nil {
		errMsg := fmt.Sprintf("Неправильный формат даты (%s). Требуемый формат: dd.mm.yyyy", songView.Text)
		log.Debugf("%s. %s. %v", generalLog, errMsg, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return 0, err
	}

	// Конвертация в songDto
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
		errMsg := fmt.Sprintf("Ошибка создания песни (%s-%s).", songDto.Group, songDto.Name)
		log.Debugf("%s. %s. %v", generalLog, errMsg, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	return id, nil
}

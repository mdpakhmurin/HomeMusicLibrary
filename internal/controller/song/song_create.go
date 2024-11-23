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
// @Description Создает новую песню на основе входных данных
// @Accept json
// @Produce json
// @Param input body SongCreate true "Входные данные для создания песни"
// @Success 200 {string} string "ID созданной песни"
// @Failure 400 {object} string "Описание ошибки"
// @Router /song [post]
func (controller *SongController) SongCreate(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s", c.Request.URL.Path)

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
	log.Infof("%s. Успешное добавление", generalLog)
}

// Создание песни с помощью сервиса
func songCreate(c *gin.Context, songView *SongCreateViewIn) (id int, err error) {
	// Получение даты
	releaseDate, err := time.Parse("02.01.2006", songView.ReleaseDate)
	if err != nil {
		errMsg := fmt.Sprintf("Неправильный формат даты (%s): %v", songView.Text, err)
		log.Debugf("%s. %s", c.Request.URL.Path, errMsg)
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
		errMsg := fmt.Sprintf("Ошибка создания песни (%s-%s): %v", songDto.Group, songDto.Name, err)
		log.Debugf("%s. %s", c.Request.URL.Path, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	return id, nil
}

package song

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Контроллер песен
type SongController struct {
}

// Общее описание обрабатываемого запроса.
func getGeneralRequestInfo(c *gin.Context) (info string) {
	return fmt.Sprintf("Запрос %s [%s]", c.Request.URL.Path, c.Request.Method)
}

// Получение данных из тела запроса.
// Сохраняет ошибку в ответ запроса.
func getBodyData(c *gin.Context, resultOut any) (err error) {
	if err = c.ShouldBindJSON(resultOut); err != nil {
		log.Debugf("%s. Ошибка получения данных из тела запроса: %v", getGeneralRequestInfo(c), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения данных из тела запроса"})
		return err
	}

	return nil
}

// Получение данных из параметров запроса.
// Сохраняет ошибку в ответ запроса.
func getRequestData(c *gin.Context, resultOut any) (err error) {
	if err = c.ShouldBind(resultOut); err != nil {
		log.Debugf("%s. Ошибка получения данных из тела запроса: %v", getGeneralRequestInfo(c), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения данных из тела запроса"})
		return err
	}

	return nil
}

// Получение даты из строки.
// Сохраняет ошибку в ответ запроса.
func parseDate(c *gin.Context, timeStr string) (date *time.Time, err error) {
	if timeStr == "" {
		return nil, nil
	}

	// Парсинг даты
	dateParseResult, err := time.Parse("02.01.2006", timeStr)
	if err != nil {
		errMsg := fmt.Sprintf("Неправильный формат даты (%s). Требуемый формат: dd.mm.yyyy", timeStr)
		log.Debugf("%s. %s. %v", getGeneralRequestInfo(c), errMsg, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return nil, err
	}

	return &dateParseResult, nil
}

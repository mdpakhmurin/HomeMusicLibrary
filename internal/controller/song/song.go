package song

import (
	"fmt"
	"net/http"

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
func getBodyData(c *gin.Context, resultOut any) (err error) {
	generalLog := fmt.Sprintf("Запрос %s", c.Request.URL.Path)
	if err = c.ShouldBindJSON(resultOut); err != nil {
		errMsg := fmt.Sprintf("%s. Ошибка получения данных из тела запроса", generalLog)
		log.Debugf("%s. %s. %v", c.Request.URL.Path, errMsg, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return err
	}

	return nil
}

// Получение данных из параметров запроса.
func getRequestData(c *gin.Context, resultOut any) (err error) {
	generalLog := fmt.Sprintf("Запрос %s", c.Request.URL.Path)
	if err = c.ShouldBind(resultOut); err != nil {
		errMsg := fmt.Sprintf("%s. Ошибка получения данных из запроса", generalLog)
		log.Debugf("%s. %s. %v", generalLog, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return err
	}

	return nil
}

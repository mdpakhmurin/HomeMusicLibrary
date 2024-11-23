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

// Получение данных из тела запроса
func getBodyData(c *gin.Context, resultOut any) (err error) {
	if err = c.ShouldBindJSON(resultOut); err != nil {
		errMsg := fmt.Sprintf("Ошибка получения данных из запроса: %v", err)
		log.Debugf("%s. %s", c.Request.URL.Path, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return err
	}

	return nil
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SongController struct {
}

// SongInfo "Получение песни по группе и названию песни"
// @Summary Получение информации о песне
// @Description Получение информации о песне по заданной группе и названию
// @Produce json
// @Param song query SongInfoDtoIn true "Данные о песне"
// @Success 200 {object} SongDetail
// @Router /song/info [get]
func (controller *SongController) GetSong(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных песни
	var input SongInfoDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		return
	}

	// Создание структуры песни с данными
	song := SongInfoDtoOut{
		Text:        "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Link:        "https://2323",
		ReleaseDate: "02.08.2009",
	}

	// Преобразование структуры песни в формат JSON
	songJSON, err := json.Marshal(song)
	if err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке конвертации
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка конвертации ответа в JSON"})
		log.Debugf("%s. Ошибка конвертации ответа (%v): %v", generalLog, song, err)
		return
	}

	// Отправка данных JSON в ответе
	c.Data(http.StatusOK, "application/json", songJSON)
	log.Infof("%s. Успешный ответ: %v", generalLog, string(songJSON))
}

// SearchSong поиск песни.
// @Summary Поиск песни
// @Description Поиск песни по заданным параметрам
// @Produce json
// @Param song query SongSearchFilter true "Данные о песне"
// @Success 200 {array} SongDescription
// @Router /song/search [get]
func (controller *SongController) SearchSong(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных поиска песни
	var input SongSearchDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		return
	}

	// Создание списка структур песен с данными
	songs := []SongSearchDtoOut{
		{Group: "cool group", Name: "super song", ReleaseDate: "23.11.2024", Link: "https://"},
		{Group: "nice gr", Name: "ok song", ReleaseDate: "10.11.2024", Link: "https://new"},
	}

	// Преобразование списка структур песен в формат JSON
	songsJSON, err := json.Marshal(songs)
	if err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке конвертации
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при конвертации в JSON"})
		log.Debugf("%s. Ошибка конвертации ответа (%v): %v", generalLog, songs, err)
		return
	}

	// Отправка данных JSON в ответе
	c.Data(http.StatusOK, "application/json", songsJSON)
	log.Infof("%s. Успешный ответ: %v", generalLog, string(songsJSON))
}

// DeleteSong удаляет песню по указанным параметрам.
// @Summary Удалить песню
// @Description Удаляет песню по указанным параметрам: group и name.
// @Produce json
// @Param song query SongRemoveDtoIn true "Данные о песне"
// @Success 200 {object} string "Сообщение об успешном удалении песни"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Router /song/ [delete]
func (controller *SongController) DeleteSong(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных удаления песни
	var input SongRemoveDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Формирование строки с сообщением об успешном удалении песни
	var responseString = fmt.Sprintf("Песня (%s: %s) успешно удалена", input.Group, input.Name)

	// Отправка ответа с сообщением об успешном удалении
	c.JSON(http.StatusOK, gin.H{"message": responseString})
	log.Infof("%s. Успешное удаление: %v", generalLog, responseString)
}

// @Summary Обновляет песню с заданным названием и группой
// @Description Добавляет новую песню
// @Accept json
// @Produce json
// @Param song body SongUpdate true "Данные о песне в формате JSON"
// @Success 200 {string} string "Сообщение о успешном добавлении песни"
// @Failure 400 {object} string "Ошибка при обработке запроса"
// @Router /song/ [put]
func (controller *SongController) SongUpdate(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s", c.Request.URL.Path)

	// Инициализация переменной для входных данных обновления песни
	var input SongUpdateDtoIn

	// Привязка входных данных к переменной input из JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		// В случае ошибки отправить ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		return
	}

	// Обновление общего лога с полученными данными
	generalLog = fmt.Sprintf("%s. %v", generalLog, input)
	log.Info(generalLog)

	// Формирование ответного сообщения об успешном обновлении песни
	answer := fmt.Sprintf("Песня (%s: %s) успешно обновлена", input.Group, input.Name)

	// Отправка ответа с сообщением об успешном обновлении
	c.JSON(http.StatusOK, gin.H{"message": answer})
	log.Infof("%s. Успешное обновление", generalLog)
}

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

// SongVersesGet получение текста песни.
// @Summary Получение текста песни
// @Description Получение куплетов песни по заданным параметрам
// @Produce json
// @Param song query SongGetVersesDtoIn true "Данные о песне"
// @Success 200 {array} string
// @Router /song/verses [get]
func (controller *SongController) SongVersesGet(c *gin.Context) {
	// Создание общего лога для запроса
	generalLog := fmt.Sprintf("Запрос %s %v", c.Request.URL.Path, c.Request.URL.Query())
	log.Infof(generalLog)

	// Инициализация переменной для входных данных получения куплетов песни
	var input SongGetVersesDtoIn

	// Проверка и привязка входных данных к переменной input
	if err := c.ShouldBind(&input); err != nil {
		// Обработка ошибки и отправка ответа с кодом ошибки
		log.Debugf("%s. Ошибка при получении данных из запроса: %v", generalLog, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Содержание куплетов песни
	verses := []string{
		"Первый куплет текста песни",
		"Припев текста песни",
		"Второй куплет текста песни",
	}

	// Отправка ответа с куплетами песни
	c.JSON(http.StatusOK, gin.H{"verses": verses})
	log.Infof("%s. Успешный ответ: %v", generalLog, verses)
}

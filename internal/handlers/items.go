package handlers

import (
	"Jokaru-py/managingEntities/models"
	"Jokaru-py/managingEntities/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Создание объекта
func (h *Handler) CreateObject(c echo.Context) error {
	req := models.ObjectRequest{}
	err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// ID из токена
	idUser := c.Get("user").(uint)

	// Проверить есть ли такой объект
	obj, err := h.connStore.GetObjectByName(&models.Object{Name: req.Name})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if obj != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(errors.New("объект с таким именем уже есть")))
	}

	// Создать
	err = h.connStore.CreateObject(&models.Object{Name: req.Name, UserID: idUser})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	// Получить
	obj, err = h.connStore.GetObjectByName(&models.Object{Name: req.Name, UserID: idUser})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	// Отдать
	result := models.ObjectResponse{
		ID:   obj.ID,
		Name: obj.Name,
	}

	return c.JSON(http.StatusOK, result)
}

// Удалить объекта
// TODO: проверку пользователя
func (h *Handler) DeleteObject(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	// Проверить есть ли такой объект
	obj, err := h.connStore.GetObjectByID(&models.Object{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if obj == nil {
		return c.JSON(http.StatusNotFound, utils.NewError(errors.New("объект с таким ID не найден")))
	}

	err = h.connStore.DeleteObjectByID(&models.Object{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, nil)
}

// Получить все объекты пользователя
func (h *Handler) GetObject(c echo.Context) error {

	// ID из токена
	idUser := c.Get("user").(uint)

	res, err := h.connStore.GetAllObjectByID(&models.Object{UserID: idUser})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, res)
}

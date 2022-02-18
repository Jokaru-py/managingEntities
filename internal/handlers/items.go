package handlers

import (
	"Jokaru-py/managingEntities/models"
	"Jokaru-py/managingEntities/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Создание объекта
func (h *Handler) CreateObject(c echo.Context) error {
	req := models.ObjectRequest{}
	err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	err = h.connStore.CreateObject(&models.ObjectDB{Name: req.Name})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, nil)
}

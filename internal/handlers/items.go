package handlers

import (
	"Jokaru-py/managingEntities/models"
	"Jokaru-py/managingEntities/pkg/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	// obj, err := h.connStore.GetObjectByName(&models.Object{Name: req.Name})
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	// }
	// if obj != nil {
	// 	return c.JSON(http.StatusInternalServerError, utils.NewError(errors.New("объект с таким именем уже есть")))
	// }

	// Создать
	obj, err := h.connStore.CreateObject(&models.Object{Name: req.Name, UserID: idUser})
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
func (h *Handler) DeleteObject(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	// ID из токена
	idUser := c.Get("user").(uint)

	// Проверить есть ли такой объект
	obj, err := h.connStore.GetObjectByID(&models.Object{Model: gorm.Model{ID: uint(id)}, UserID: idUser})
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

	if len(res) == 0 {
		return c.JSON(http.StatusNotFound, utils.NewError(errors.New("ничего не найдено")))
	}

	return c.JSON(http.StatusOK, res)
}

// Передать один объект другому
func (h *Handler) SendObject(c echo.Context) error {

	req := models.SendObjectRequest{}
	err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// ID из токена
	idUser := c.Get("user").(uint)

	// Проверить есть ли такой объект
	obj, err := h.connStore.GetObjectByID(&models.Object{Model: gorm.Model{ID: req.ID}, UserID: idUser})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if obj == nil {
		return c.JSON(http.StatusNotFound, utils.NewError(errors.New("объект с таким ID не найден")))
	}

	// Проверить нового владельца
	newUser, err := h.connStore.GetUser(req.Login)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewError(errors.New("пользователь с таким логином не найден")))
	}
	if newUser.Pass == "" {
		return c.JSON(http.StatusNotFound, utils.NewError(errors.New("пользователь с таким логином не найден")))
	}

	// Поменять владельца объекта
	err = h.connStore.UpdateObject(&models.Object{Model: gorm.Model{ID: obj.ID}, UserID: newUser.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	// Отправить ссылку
	var link = fmt.Sprintf("localhost:%s/api/", os.Getenv("PORT"))

	// texthash := utils.GenerateLink(req.Login, req.ID)
	texthash := utils.EndcodeLink(fmt.Sprintf("login=%s id=%d", req.Login, req.ID))

	return c.JSON(http.StatusOK, fmt.Sprintf("%s%s", link, texthash))
}

// Получить переданный объект
func (h *Handler) GetNewObject(c echo.Context) error {
	query := utils.DecodeLink(c.Param("get"))
	log.Println(utils.DecodeLink(query))

	// login := strings.ReplaceAll(strings.Split(query, " ")[0], "login=", "")
	idObject := strings.ReplaceAll(strings.Split(query, " ")[1], "id=", "")
	idObjectInt, err := strconv.Atoi(idObject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	// ID из токена
	idUser := c.Get("user").(uint)

	// Проверка ID пользователя из токена и пути
	// user, err := h.connStore.GetUser(login)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	// }
	// if user.ID != id {
	// 	return c.JSON(http.StatusInternalServerError, utils.NewError(errors.New("нет доступа")))
	// }

	// Проверить есть ли такой объект
	obj, err := h.connStore.GetObjectByID(&models.Object{Model: gorm.Model{ID: uint(idObjectInt)}, UserID: idUser})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if obj == nil {
		return c.JSON(http.StatusNotFound, utils.NewError(errors.New("объект с таким данными не найден")))
	}

	return c.JSON(http.StatusOK, obj)
}

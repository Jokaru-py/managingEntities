package handlers

import (
	"Jokaru-py/managingEntities/models"
	"Jokaru-py/managingEntities/pkg/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Авторизация
func (h *Handler) SignIn(c echo.Context) error {
	req := &models.UserRequest{}
	err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// Проверка есть ли уже пользователь в БД
	user, err := h.connStore.GetUser(req.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if user != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(errors.New("пользователь уже существует")))
	}

	err = utils.PasswordCheck(user.Pass, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, utils.GenerateJWT(user.ID))
}

// Регистрация
func (h *Handler) SignUp(c echo.Context) error {
	req := &models.UserRequest{}
	err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// Проверка есть ли уже пользователь в БД
	user, err := h.connStore.GetUser(req.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if user != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(errors.New("пользователь уже существует")))
	}

	// Создать пользователя
	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	err = h.connStore.CreateUser(&models.UsersDB{Login: req.Login, Pass: hashPass})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, nil)
}

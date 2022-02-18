package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	v1.POST("registration", h.SignUp) // Регистрация
	v1.POST("login", h.SignIn)        // Авторизация
}

package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	v1.POST("/registration", h.SignUp) // Регистрация
	v1.POST("/login", h.SignIn)        // Авторизация

	// Работа с объектами
	items := v1.Group("/items")
	items.POST("/new", h.CreateObject)   // Создать объект
	items.DELETE("/:id", h.DeleteObject) // Удалить объект
	items.GET("", h.GetObject)           // Получить объект
}

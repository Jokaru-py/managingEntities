package handlers

import (
	"Jokaru-py/managingEntities/pkg/router/middleware"
	"Jokaru-py/managingEntities/pkg/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {

	// Auth
	v1.POST("/registration", h.SignUp) // Регистрация
	v1.POST("/login", h.SignIn)        // Авторизация

	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	// Работа с объектами
	items := v1.Group("/items", jwtMiddleware)
	items.POST("/new", h.CreateObject)   // Создать объект
	items.DELETE("/:id", h.DeleteObject) // Удалить объект
	items.GET("", h.GetObject)           // Получить объект

	//
	v1.POST("/send", h.SendObject, jwtMiddleware)  // Генерация ссылки для передачи объекта
	v1.GET("/:get", h.GetNewObject, jwtMiddleware) // Переход по ссылке для получения объекта
}

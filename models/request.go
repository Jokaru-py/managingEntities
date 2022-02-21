package models

import "github.com/labstack/echo/v4"

type UserRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *UserRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type ObjectRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *ObjectRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type SendObjectRequest struct {
	ID    uint   `json:"id" validate:"required"`    // id передаваемого объекта
	Login string `json:"login" validate:"required"` // логин принимающего пользователя
}

func (r *SendObjectRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

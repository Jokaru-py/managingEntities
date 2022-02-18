package handlers

import (
	"Jokaru-py/managingEntities/internal/store"
)

type Handler struct {
	connStore store.ConnStore
}

func NewHandler(conn store.ConnStore) *Handler {
	return &Handler{
		connStore: conn,
	}
}

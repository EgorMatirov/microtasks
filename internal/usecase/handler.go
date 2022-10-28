package usecase

import (
	"github.com/EgorMatirov/microtasks/internal/domain"
)

// сюда из интерактора!
type CRUD interface {
}

type Handler interface {
	CRUD
}

type HandlerConstructor struct {
	Crud domain.Repository
}

func (c *HandlerConstructor) New() Handler {
	return &interactor{
		crudRepo: c.Crud,
	}
}

package usecase

import (
	"github.com/EgorMatirov/microtasks/internal/domain"
)

type interactor struct {
	crudRepo domain.Repository
}

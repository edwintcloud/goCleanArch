package usecases

import (
	"goCleanArch/repositories"
)

// Usecase struct
type Usecase struct {
	repo repositories.Repository
}

// NewUsecase creates a new usecase
func NewUsecase(r repositories.Repository) *Usecase {
	return &Usecase{
		repo: r,
	}
}

// FindByID finds a resource by id
func (u *Usecase) FindByID(id interface{}) (interface{}, error) {
	return u.repo.FindByID(id)
}

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

// FindAll finds all resources for usecase
func (u *Usecase) FindAll() ([]interface{}, error) {
	return u.repo.FindAll()
}

// Create creates a new resource with specified data and returns new resource
func (u *Usecase) Create(id interface{}, data interface{}) (interface{}, error) {

	// Insert resource into repo
	if err := u.repo.Insert(data); err != nil {
		return nil, err
	}

	// find resource by id and return
	return u.repo.FindByID(id)

}

// DeleteByID deletes a resource by id
func (u *Usecase) DeleteByID(id interface{}) error {
	return u.repo.DeleteByID(id)
}

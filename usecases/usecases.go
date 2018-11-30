package usecases

import (
	"errors"
	"goCleanArch/repositories"

	"golang.org/x/crypto/bcrypt"
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

// UpdateByID updates a resource by id
func (u *Usecase) UpdateByID(id interface{}, updates interface{}) error {

	// Update resource by id
	if err := u.repo.UpdateByID(id, updates); err != nil {
		return err
	}

	// return nil if no error
	return nil
}

// DeleteByID deletes a resource by id
func (u *Usecase) DeleteByID(id interface{}) error {
	return u.repo.DeleteByID(id)
}

// Login find a user by query and compares given password with repo password
func (u *Usecase) Login(query interface{}, password interface{}) error {

	// find user by email
	result, err := u.repo.FindByQuery(query)
	if err != nil {
		return err
	}

	// if len result == 0, return user not found
	if len(result) == 0 {
		return errors.New("user not found")
	}

	// compare passwords and return result
	if err := bcrypt.CompareHashAndPassword([]byte(result[0]["password"].(string)), []byte(password.(string))); err != nil {
		return errors.New("invalid password")
	}
	return nil

}

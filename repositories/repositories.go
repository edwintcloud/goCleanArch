package repositories

import (
	"goCleanArch/models"

	"github.com/juju/mgosession"
)

// Repository interface definition
type Repository interface {
	FindByID(id string) (*models.User, error)
	FindAll() ([]*models.User, error)
	Create(user *models.User) (*models.User, error)
	UpdateByID(id string, updates map[string]interface{}) (*models.User, error)
	DeleteByID(id string) error
}

// mongoRepository struct representing database connection for internal use
type mongoRepository struct {
	pool *mgosession.Pool
}

// NewMongoRepository sets mongoRepository pool connection
func NewMongoRepository(p *mgosession.Pool) Repository {
	return &mongoRepository{
		pool: p,
	}
}

func (r *mongoRepository) FindByID(id string) (*models.User, error) {

}

func (r *mongoRepository) FindAll() ([]*models.User, error) {

}

func (r *mongoRepository) Create(user *models.User) (*models.User, error) {

}

func (r *mongoRepository) UpdateByID(id string, updates map[string]interface{}) (*models.User, error) {

}

func (r *mongoRepository) DeleteByID(id string) error {

}

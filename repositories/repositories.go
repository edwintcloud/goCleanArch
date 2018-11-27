package repositories

import (
	"github.com/juju/mgosession"
)

// Repository interface definition
type Repository interface {
	FindByID(id interface{}) (*interface{}, error)
	FindAll() ([]*interface{}, error)
	Create(user *interface{}) (*interface{}, error)
	UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error)
	DeleteByID(id interface{}) error
}

// mongoRepository struct representing database connection for internal use
type mongoRepository struct {
	pool  *mgosession.Pool
	model interface{}
}

// NewMongoRepository sets mongoRepository pool connection m is the model
func NewMongoRepository(p *mgosession.Pool, m interface{}) Repository {
	return &mongoRepository{
		pool:  p,
		model: m,
	}
}

func (r *mongoRepository) FindByID(id interface{}) (*interface{}, error) {
	// result := models.User{}
	// get interface type name and use as collection name, db name will come from viper config
	session := r.pool.Session(nil)

}

func (r *mongoRepository) FindAll() ([]*interface{}, error) {

}

func (r *mongoRepository) Create(user *interface{}) (*interface{}, error) {

}

func (r *mongoRepository) UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error) {

}

func (r *mongoRepository) DeleteByID(id interface{}) error {

}

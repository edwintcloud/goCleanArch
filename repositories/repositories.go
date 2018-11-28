package repositories

import (
	"reflect"

	"github.com/juju/mgosession"
	"github.com/spf13/viper"
)

// Repository interface definition
type Repository interface {
	FindByID(id interface{}) (*interface{}, error)
	// FindAll() ([]*interface{}, error)
	// Create(user *interface{}) (*interface{}, error)
	// UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error)
	// DeleteByID(id interface{}) error
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
	var result interface{}

	// get interface type name and use as collection name, db name will come from viper config
	session := r.pool.Session(nil)

	// get collection name
	collection := getType(r.model)

	// get database name
	db := viper.GetString("database.name")

	// create collection instance
	c := session.DB(db).C(collection)

	// find by id or return nil and err
	if err := c.FindId(id).One(&result); err != nil {
		return nil, err
	}

	// otherwise return result and no error
	return &result, nil
}

// func (r *mongoRepository) FindAll() ([]*interface{}, error) {

// }

// func (r *mongoRepository) Create(user *interface{}) (*interface{}, error) {

// }

// func (r *mongoRepository) UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error) {

// }

// func (r *mongoRepository) DeleteByID(id interface{}) error {

// }

// returns string representation of interface{} type
func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

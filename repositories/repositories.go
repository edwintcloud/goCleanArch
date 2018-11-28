package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/viper"
)

// Repository interface definition
type Repository interface {
	FindByID(id interface{}) (interface{}, error)
	// FindAll() ([]*interface{}, error)
	// Create(user *interface{}) (*interface{}, error)
	// UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error)
	// DeleteByID(id interface{}) error
}

// mongoRepository struct representing database connection for internal use
type mongoRepository struct {
	db         *mgo.Session
	collection string
}

// NewMongoRepository sets mongoRepository pool connection m is the model
func NewMongoRepository(d *mgo.Session, c string) Repository {
	return &mongoRepository{
		db:         d,
		collection: c,
	}
}

func (r *mongoRepository) FindByID(id interface{}) (interface{}, error) {
	var result interface{}

	// get database name
	db := viper.GetString("database.name")

	// create collection instance
	c := r.db.DB(db).C(r.collection)

	// find by id or return nil and err
	if err := c.FindId(bson.ObjectIdHex(id.(string))).One(&result); err != nil {
		return nil, err
	}

	// otherwise return result and no error
	return result, nil
}

// func (r *mongoRepository) FindAll() ([]*interface{}, error) {

// }

// func (r *mongoRepository) Create(user *interface{}) (*interface{}, error) {

// }

// func (r *mongoRepository) UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error) {

// }

// func (r *mongoRepository) DeleteByID(id interface{}) error {

// }

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
	Insert(data interface{}) error
	// UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error)
	// DeleteByID(id interface{}) error
}

// mongoRepository struct representing database connection for internal use
type mongoRepository struct {
	c *mgo.Collection
}

// NewMongoRepository sets mongoRepository pool connection m is the model
func NewMongoRepository(d *mgo.Session, c string) Repository {
	// get database name from config
	db := viper.GetString("database.name")

	return &mongoRepository{
		c: d.DB(db).C(c),
	}
}

func (r *mongoRepository) FindByID(id interface{}) (interface{}, error) {
	var result interface{}

	// find by id or return nil and err
	if err := r.c.FindId(bson.ObjectIdHex(id.(string))).One(&result); err != nil {
		return nil, err
	}

	// otherwise return result and no error
	return result, nil
}

// func (r *mongoRepository) FindAll() ([]*interface{}, error) {

// }

func (r *mongoRepository) Insert(data interface{}) error {

	// create new document
	if err := r.c.Insert(data); err != nil {
		return err
	}

	// return all good
	return nil

}

// func (r *mongoRepository) UpdateByID(id interface{}, updates map[string]interface{}) (*interface{}, error) {

// }

// func (r *mongoRepository) DeleteByID(id interface{}) error {

// }

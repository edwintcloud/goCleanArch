package main

import (
	"fmt"
	mongo_controllers "goCleanArch/controllers/mongoControllers"
	"goCleanArch/models"
	"goCleanArch/repositories"
	"goCleanArch/usecases"

	"net/http"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	// set config file for viper to use as config.json
	viper.SetConfigFile(`config.json`)
	// read config from file and panic on error
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	// create new instance of echo http server
	e := echo.New()

	// register controllers based on database config
	switch db := viper.GetString("database.type"); db {
	case "mongo":
		session := initializeMongo(e)
		defer session.Close()
	default:
		panic("No database type specified or specified type not implemented!")
	}

	// catch all route
	e.Any("*", func(c echo.Context) error {
		err := fmt.Sprintf("Bad Request - %s %s", c.Request().Method, c.Request().RequestURI)
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err})
	})

	// start echo server, panic on failure
	if err := e.Start(viper.GetString("server.address")); err != nil {
		panic(err)
	}
}

func initializeMongo(e *echo.Echo) *mgo.Session {
	// connect to db based on config
	session, err := mgo.Dial(viper.GetString("database.uri"))
	if err != nil {
		panic(err)
	}

	// create new repo based on config, passing in pool and model
	userRepository := repositories.NewMongoRepository(session, "users")

	// pass repo to usecase
	userUsecase := usecases.NewUsecase(userRepository)

	// pass echo and usecase to controller, registering routes
	mongo_controllers.InitUsers(e, userUsecase)

	// return the session to main so we can defer session close
	return session
}

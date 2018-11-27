package main

import (
	"net/http"

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

	// test json message
	message := map[string]string{
		"message": "Welcome to my api!",
	}

	// test route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, message)
	})

	// start echo server, panic on failure
	if err := e.Start(viper.GetString("server.address")); err != nil {
		panic(err)
	}
}

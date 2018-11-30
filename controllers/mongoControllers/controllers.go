package controllers

import (
	"goCleanArch/models"
	"goCleanArch/usecases"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

type handler struct {
	Usecase *usecases.Usecase
}

// InitUsers initializes user controller and routes
func InitUsers(e *echo.Echo, u *usecases.Usecase) {
	handler := &handler{
		Usecase: u,
	}

	routes := e.Group("/api/v1/users")
	{
		routes.GET("", handler.getAllUsers)
		routes.GET("/:id", handler.getUserByID)
		routes.POST("", handler.createUser)
	}
}

func (h *handler) getUserByID(c echo.Context) error {
	id := c.Param("id")

	// make sure id is not empty
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: "No Id param specified"})
	}

	// make sure id is valid object id
	if !bson.IsObjectIdHex(id) {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: "invalid objectid"})
	}

	// find by id
	result, err := h.Usecase.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return result
	return c.JSON(http.StatusOK, result)

}

func (h *handler) getAllUsers(c echo.Context) error {

	// find all
	result, err := h.Usecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return results
	return c.JSON(http.StatusOK, result)

}

func (h *handler) createUser(c echo.Context) error {
	id := bson.NewObjectId()
	curTime := time.Now()
	data := models.User{
		ID:        id,
		UpdatedAt: curTime,
		CreatedAt: curTime,
	}

	// bind request body to user struct
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// create new user in db
	result, err := h.Usecase.Create(id.Hex(), data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return result
	return c.JSON(http.StatusOK, result)

}

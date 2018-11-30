package users

import (
	"fmt"
	"goCleanArch/models"
	"goCleanArch/usecases"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	Usecase *usecases.Usecase
}

// Init initializes user controller and routes
func Init(e *echo.Echo, u *usecases.Usecase) {
	handler := &handler{
		Usecase: u,
	}

	routes := e.Group("/api/v1/users")
	{
		routes.GET("", handler.getAllUsers)
		routes.GET("", handler.getAllUsers)
		routes.GET("/:id", handler.getUserByID)
		routes.POST("", handler.createUser)
		routes.DELETE("/:id", handler.deleteUserByID)
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
	user := models.User{
		ID:        id,
		UpdatedAt: curTime,
		CreatedAt: curTime,
	}

	// bind request body to user struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	// create new user in db
	result, err := h.Usecase.Create(id.Hex(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return result
	return c.JSON(http.StatusOK, result)

}

func (h *handler) deleteUserByID(c echo.Context) error {
	id := c.Param("id")

	// make sure id is not empty
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: "No Id param specified"})
	}

	// make sure id is valid object id
	if !bson.IsObjectIdHex(id) {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: "invalid objectid"})
	}

	// delete by id
	err := h.Usecase.DeleteByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return result
	result := models.ResponseMsg{Message: fmt.Sprintf("User %s deleted successfully", id)}
	return c.JSON(http.StatusOK, result)

}

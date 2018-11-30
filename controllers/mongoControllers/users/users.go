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
		routes.POST("/login", handler.loginUser)
		routes.PUT("/:id", handler.updateUserByID)
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

func (h *handler) updateUserByID(c echo.Context) error {
	id := c.Param("id")
	updates := bson.M{}

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

	// Bind req body to bson map
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// convert result to bson map
	resultBson := result.(bson.M)

	// make changes to resultBson based on req body
	for k := range updates {
		if k == "password" {
			// hash new password
			hash, _ := bcrypt.GenerateFromPassword([]byte(updates[k].(string)), bcrypt.DefaultCost)
			resultBson[k] = string(hash)
		} else if k == "updated_at" {
			resultBson[k] = time.Now()
		} else {
			resultBson[k] = updates[k]
		}
	}

	// update by id
	err = h.Usecase.UpdateByID(id, resultBson)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return updated user
	return c.JSON(http.StatusOK, resultBson)

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

func (h *handler) loginUser(c echo.Context) error {
	user := make(map[string]interface{})

	// bind req body to bson map
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// make sure bson map has email
	if _, ok := user["email"]; !ok {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: "missing email field"})
	}

	// make sure bson map has password
	if _, ok := user["password"]; !ok {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: "missing password field"})
	}

	// login user
	if err := h.Usecase.Login(bson.M{"email": user["email"]}, user["password"]); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
	}

	// return result
	result := models.ResponseMsg{Message: fmt.Sprintf("User %s logged in successfully", user["email"])}
	return c.JSON(http.StatusOK, result)

}

package mongoControllers

import (
	"goCleanArch/models"
	"goCleanArch/usecases"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Error string `json:"error"`
}

// Handler represent the httphandler for article
type Handler struct {
	Usecase *usecases.Usecase
}

// InitUsers initializes user controller and routes
func InitUsers(e *echo.Echo, u *usecases.Usecase) {
	handler := &Handler{
		Usecase: u,
	}

	routes := e.Group("/api/v1/users")
	{
		routes.GET("/:id", handler.getUserByID)
	}
}

func (h *Handler) getUserByID(c echo.Context) error {
	id := c.Param("id")
	result := models.User{}

	// make sure id is not empty
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, ResponseError{Error: "No Id param specified"})
	}

	// make sure id is valid object id
	if !bson.IsObjectIdHex(id) {
		return c.JSON(http.StatusBadRequest, ResponseError{Error: "invalid objectid"})
	}

	// find by id
	foundResult, err := h.Usecase.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
	}

	// convert to struct and return result
	bsonBytes, _ := bson.Marshal(foundResult)
	bson.Unmarshal(bsonBytes, &result)
	return c.JSON(http.StatusOK, result)

}

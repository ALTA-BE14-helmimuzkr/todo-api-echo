package controller

import (
	"net/http"

	"github.com/ALTA-BE14-helmimuzkr/model"
	"github.com/labstack/echo"
)

type UserController struct {
	model *model.UserModel
}

func (controller *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}

		user, err := controller.model.Insert(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada server",
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    nil,
			"message": "Kesalahan pada server",
		})
	}
}

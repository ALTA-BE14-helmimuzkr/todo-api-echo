package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ALTA-BE14-helmimuzkr/model"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Model *model.UserModel
}

func (controller *UserController) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Parsing data ke user
		user := model.User{}
		if err := ctx.Bind(&user); err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}

		// Hashing password menggunakan bcrypt
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada server",
			})
		}
		// Assign hasil hash yang sudah diconvert ke string ke dalam Password
		user.Password = string(hash)

		// Insert data user ke dalam database
		user, err = controller.Model.Insert(user)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada server",
			})
		}

		// Return response ketika  sukses
		return ctx.JSON(http.StatusCreated, map[string]interface{}{
			"data":    nil,
			"message": "Berhasil melakukan registrasi",
		})
	}
}

func (controller *UserController) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Parsing from request body
		user := model.User{}
		if err := ctx.Bind(&user); err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}
		// Parsing from request parameter
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}
		// Assign id dari parameter ke user.ID
		user.ID = uint(id)

		// Hashing password menggunakan bcrypt
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada server",
			})
		}
		// Assign hasil hash yang sudah diconvert ke string ke dalam Password
		user.Password = string(hash)

		// Update
		user, err = controller.Model.Update(user)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada server",
			})
		}

		// Success
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data":    nil,
			"message": "Berhasil update data user",
		})
	}
}

func (controller *UserController) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}

		err = controller.Model.Delete(id)
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada sisi server",
			})
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data":    nil,
			"message": "Berhasil hapus data user",
		})
	}
}

func (controller *UserController) GetUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}

		user, err := controller.Model.FindById(id)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada sisi server",
			})
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil mendapatkan data user",
			"data": model.UserResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Telephone: user.Telephone,
			},
		})
	}
}

func (controller *UserController) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		users, err := controller.Model.FindAll()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada sisi server",
			})
		}

		userResponses := []model.UserResponse{}
		for _, v := range users {
			user := model.UserResponse{}
			user.ID = v.ID
			user.Name = v.Name
			user.Email = v.Email
			user.Telephone = v.Telephone

			userResponses = append(userResponses, user)
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil mendapatkan data user",
			"data":    userResponses,
		})
	}
}

func (controller *UserController) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		inputUser := model.User{}
		if err := ctx.Bind(&inputUser); err != nil {
			log.Println(err.Error())
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}

		user, err := controller.Model.FindByEmail(inputUser.Email)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Kesalahan pada sisi server",
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Kesalahan input client",
			})
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil melakukan login",
			"data": model.UserResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Telephone: user.Telephone,
			},
		})
	}
}

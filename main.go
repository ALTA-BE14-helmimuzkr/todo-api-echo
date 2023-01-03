package main

import (
	"fmt"
	"log"

	"github.com/ALTA-BE14-helmimuzkr/config"
	"github.com/ALTA-BE14-helmimuzkr/controller"
	"github.com/ALTA-BE14-helmimuzkr/database"
	"github.com/ALTA-BE14-helmimuzkr/model"
	"github.com/labstack/echo"
)

func main() {
	// Setup Echo
	e := echo.New()

	// Setup Config
	c := config.InitConfig()
	fmt.Println(c)

	// Setup Dataabse
	db := database.OpenConnectionMysql(c)

	// Setup Migration
	model.Migrate(db)

	// Setup Repository
	userModel := model.UserModel{DB: db}

	// Setup Controller
	userController := controller.UserController{Model: &userModel}

	// Router
	e.POST("/users", userController.Create())
	e.PUT("/users/:id", userController.Update())
	e.DELETE("/users/:id", userController.Delete())
	e.GET("/users/:id", userController.GetUser())
	e.GET("/users", userController.GetAll())
	e.GET("/users/login", userController.Login())

	// Server
	if err := e.Start(":8080"); err != nil {
		log.Println("RUNNING SERVER FAILED", err)
		return
	}
}

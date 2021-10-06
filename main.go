package main

import (
	"github.com/darienkentanu/API-CRUD-User-Using-Database/model"
	"github.com/darienkentanu/API-CRUD-User-Using-Database/controller"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)



func main() {
	// create a new echo instance
	model.InitDB()
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

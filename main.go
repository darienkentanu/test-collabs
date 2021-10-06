package main

import (
	"net/http"
	"strconv"

	"github.com/darienkentanu/API-CRUD-User-Using-Database/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type M map[string]interface{}

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User
	if err := model.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, M{"message": "success get all users",
		"users": users})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "id not found")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	return c.JSON(http.StatusOK, user)
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "failed create new user")
	}
	if err := model.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, M{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "id not found")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	if err := model.DB.Delete(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"message": "please input a valid id"})
	}
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := model.DB.Save(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	// create a new echo instance
	model.InitDB()
	e := echo.New()
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

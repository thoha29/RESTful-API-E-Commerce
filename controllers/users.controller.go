package controllers

import (
	"e-commerce/helpers"
	"e-commerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetcUsers(c echo.Context) error {
	result, err := models.FetcUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateUser(c echo.Context) error {
	user_name := c.FormValue("name")
	user_pass := c.FormValue("pass")
	user_email := c.FormValue("email")
	alamat := c.FormValue("alamat")

	result, err := models.CreateUser(user_name, user_pass, user_email, alamat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	user_id := c.Param("id")
	user_name := c.FormValue("name")
	user_pass := c.FormValue("pass")
	user_email := c.FormValue("email")
	alamat := c.FormValue("alamat")

	conv_id, err := strconv.Atoi(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(conv_id, user_name, user_pass, user_email, alamat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	user_id := c.Param("id")

	conv_id, err := strconv.Atoi(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func LoginUser(c echo.Context) error {
	user_email := c.FormValue("email")
	user_pass := c.FormValue("pass")

	res, user, err := models.LoginUser(user_email, user_pass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	tokenString, err := helpers.GenerateToken(user.User_id, user.User_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

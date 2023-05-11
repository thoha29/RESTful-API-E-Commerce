package controllers

import (
	"e-commerce/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	user_id := int(userId)

	productId, err := strconv.Atoi(c.FormValue("product_id"))
	qty, err := strconv.Atoi(c.FormValue("qty"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	err = models.AddToCart(user_id, productId, qty)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to add to cart")
	}

	return c.JSON(http.StatusOK, "Product added to cart")
}

func FetchCartByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	user_id := int(userId)

	result, err := models.FetchCartByUserId(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCart(c echo.Context) error {

	cart_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := models.DeleteCart(cart_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func EditQty(c echo.Context) error {

	cart_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	qty, err := strconv.Atoi(c.FormValue("qty"))

	if qty <= 0 {
		return c.JSON(http.StatusBadRequest, "Invalid quantity value")
	}

	res, err := models.EditQty(cart_id, qty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

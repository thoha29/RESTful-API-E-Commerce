package controllers

import (
	"e-commerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchProductsWithGenres(c echo.Context) error {
	result, err := models.FetchProductsWithGenres()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateProduct(c echo.Context) error {
	product_name := c.FormValue("product")
	price := c.FormValue("price")
	stock := c.FormValue("stock")
	genre_id := c.FormValue("genre_id")

	conv_price, err := strconv.Atoi(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_stock, err := strconv.Atoi(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_genreId, err := strconv.Atoi(genre_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.CreateProduct(product_name, conv_price, conv_stock, conv_genreId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	product_id := c.Param("id")

	conv_id, err := strconv.Atoi(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteProduct(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
	product_id := c.Param("id")
	product_name := c.FormValue("product")
	price := c.FormValue("price")
	stock := c.FormValue("stock")
	genre_id := c.FormValue("genre_id")

	conv_id, err := strconv.Atoi(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_price, err := strconv.Atoi(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_stock, err := strconv.Atoi(stock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_genreId, err := strconv.Atoi(genre_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateProduct(conv_id, product_name, conv_price, conv_stock, conv_genreId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

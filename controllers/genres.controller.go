package controllers

import (
	"e-commerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// FetchGenres function retrieves all genres
func FetchGenres(c echo.Context) error {
	result, err := models.FetchGenres()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// FetchGenre function retrieves a genre by ID
func FetchGenre(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := models.FetchGenre(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// CreateGenre function Creates a new genre
func CreateGenre(c echo.Context) error {
	genre_name := c.FormValue("genre")

	if genre_name == "" {
		return c.JSON(http.StatusBadRequest, "Data can't be empty")
	}
	result, err := models.CreateGenre(genre_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateGenre function updates a genre by ID
func UpdateGenre(c echo.Context) error {
	genre_id := c.Param("id")
	genre_name := c.FormValue("genre")

	if genre_name == "" {
		return c.JSON(http.StatusBadRequest, "Data can't be empty")
	}

	conv_id, err := strconv.Atoi(genre_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateGenre(conv_id, genre_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// DeleteGenre function deletes a genre by ID
func DeleteGenre(c echo.Context) error {
	genre_id := c.Param("id")
	conv_id, err := strconv.Atoi(genre_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteGenre(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

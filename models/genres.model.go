package models

import (
	"database/sql"
	"e-commerce/db"
	"net/http"
)

type Genres struct {
	Genre_id   int    `json:"genre_id"`
	Genre_name string `json:"genre_name"`
}

func FetchGenres() (Response, error) {
	var obj Genres
	var arrobj []Genres
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `genres`"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Genre_id, &obj.Genre_name)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func FetchGenre(genre_id int) (Response, error) {
	var obj Genres
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `genres` WHERE `genre_id` = ?"

	err := con.QueryRow(sqlStatement, genre_id).Scan(&obj.Genre_id, &obj.Genre_name)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = http.StatusNotFound
			res.Message = "Genre not found"
			return res, nil
		} else {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func CreateGenre(genre_name string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO `genres`(`genre_name`) VALUES (?)"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(genre_name)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"lastInsertedId": lastInsertedId,
	}

	return res, nil
}

func UpdateGenre(genre_id int, genre_name string) (Response, error) {
	var res Response
	con := db.CreateCon()

	// Check if genre_id exists
	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM genres WHERE genre_id = ?", genre_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Genre not found"
		return res, nil
	}

	sqlStatement := "UPDATE genres SET genre_name = ? WHERE genre_id = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(genre_name, genre_id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func DeleteGenre(genre_id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	// Check if genre_id exists
	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM genres WHERE genre_id = ?", genre_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Genres not found"
		return res, nil
	}

	sqlStatement := "DELETE FROM genres WHERE genre_id = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(genre_id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

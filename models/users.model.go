package models

import (
	"database/sql"
	"e-commerce/db"
	"e-commerce/helpers"
	"fmt"
	"net/http"
)

type Users struct {
	User_id    int    `json:"user_id"`
	User_name  string `json:"user_name"`
	User_pass  string `json:"user_pass"`
	User_email string `json:"user_email"`
	Alamat     string `json:"alamat"`
}

func FetcUsers() (Response, error) {
	var obj Users
	var arrobj []Users
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `users`"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.User_id, &obj.User_name, &obj.User_pass, &obj.User_email, &obj.Alamat)

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

func CreateUser(user_name string, user_pass string, user_email string, alamat string) (Response, error) {
	var res Response
	con := db.CreateCon()

	hashedPassword, err := helpers.HashPassword(user_pass)
	if err != nil {
		return res, err
	}

	sqlStatement := "INSERT INTO `users`(`user_name`, `user_pass`, `user_email`, `alamat`) VALUES (?, ?, ?, ?)"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(user_name, hashedPassword, user_email, alamat)
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

func UpdateUser(user_id int, user_name string, user_pass string, user_email string, alamat string) (Response, error) {
	var res Response
	con := db.CreateCon()

	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM users WHERE user_id = ?", user_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "User not found"
		return res, nil
	}

	hashedPassword, err := helpers.HashPassword(user_pass)
	if err != nil {
		return res, err
	}

	sqlStatement := "UPDATE users SET user_name = ?, user_pass = ?, user_email = ?, alamat = ? WHERE user_id = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(user_name, hashedPassword, user_email, alamat, user_id)
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

func DeleteUser(user_id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	// Check if user_id exists
	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM users WHERE user_id = ?", user_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "User not found"
		return res, nil
	}

	sqlStatement := "DELETE FROM users WHERE user_id = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(user_id)
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

func GetUserByEmail(user_email string) (Users, error) {
	var user Users
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `users` WHERE `user_email` = ? LIMIT 1"

	err := con.QueryRow(sqlStatement, user_email).Scan(&user.User_id, &user.User_name, &user.User_pass, &user.User_email, &user.Alamat)
	if err == sql.ErrNoRows {
		fmt.Println("User_email is not found")
		return user, err
	}

	if err != nil {
		fmt.Println("Query error:")
		return user, err
	}

	return user, nil
}

func LoginUser(user_email string, user_pass string) (bool, Users, error) {
	var obj Users
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `users` WHERE `user_email` = ? LIMIT 1"

	err := con.QueryRow(sqlStatement, user_email).Scan(&obj.User_id, &obj.User_name, &obj.User_pass, &obj.User_email, &obj.Alamat)
	if err == sql.ErrNoRows {
		fmt.Println("User_email is not found")
		return false, obj, err
	}

	if err != nil {
		fmt.Println("Query error:")
		return false, obj, err
	}

	passwordMatch, err := helpers.CheckPasswordHash(user_pass, obj.User_pass)
	if !passwordMatch {
		fmt.Println("Password hash error:", err)
		return false, obj, err
	}

	return true, obj, nil
}

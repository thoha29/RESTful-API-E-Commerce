package models

import (
	"database/sql"
	"e-commerce/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Products struct {
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Genre_id     int    `json:"genre_id"`
}

func FetchProducts() (Response, error) {
	var obj Products
	var arrobj []Products
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `products`"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Product_id, &obj.Product_name, &obj.Price, &obj.Stock, &obj.Genre_id)

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

func FetchProduct(product_id int) (Response, error) {
	var obj Products
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `products` WHERE `product_id` = ?"

	err := con.QueryRow(sqlStatement, product_id).Scan(&obj.Product_id, &obj.Product_name, &obj.Price, &obj.Stock, &obj.Genre_id)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = http.StatusNotFound
			res.Message = "Products not found"
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

func CreateProduct(product_name string, price int, stock int, genre_id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO `products`(`product_name`, `price`, `stock`, `genre_id`) VALUES (?, ?, ?, ?)"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(product_name, price, stock, genre_id)
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

func UpdateProduct(product_id int, product_name string, price int, stock int, genre_id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	// Check if product_id exists
	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM products WHERE product_id = ?", product_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Product not found"
		return res, nil
	}

	sqlStatement := "UPDATE products SET product_name = ?, price = ?, stock = ?, genre_id = ? WHERE product_id = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(product_name, price, stock, genre_id, product_id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	if rowsAffected == 0 {
		res.Status = http.StatusInternalServerError
		res.Message = "Failed to update product"
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func DeleteProduct(product_id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	// Check if product_id exists
	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM products WHERE product_id = ?", product_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Product not found"
		return res, nil
	}

	sqlStatement := "DELETE FROM products WHERE product_id = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(product_id)
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
		"product_id":    int64(product_id),
		"rows_affected": rowsAffected,
	}

	return res, nil
}

////////////////////////////////////////////////////////////////

type ProductWithGenre struct {
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Genre        struct {
		Genre_id   int    `json:"genre_id"`
		Genre_name string `json:"genre_name"`
	} `json:"genre"`
}

func FetchProductsWithGenres() (Response, error) {
	var obj ProductWithGenre
	var arrobj []ProductWithGenre
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT products.product_id, products.product_name, products.price, products.stock, genres.genre_id, genres.genre_name FROM `products` JOIN `genres` ON products.genre_id = genres.genre_id"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Product_id, &obj.Product_name, &obj.Price, &obj.Stock, &obj.Genre.Genre_id, &obj.Genre.Genre_name)

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

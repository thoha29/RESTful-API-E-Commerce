package models

import (
	"e-commerce/db"
	"net/http"
)

type Cart struct {
	Cart_id    int `json:"cart_id"`
	User_id    int `json:"user_id"`
	Product_id int `json:"product_id"`
	Qty        int `json:"qty"`
}

func AddToCart(user_id int, product_id int, qty int) error {
	con := db.CreateCon()

	sqlStatement := "INSERT INTO `carts` (`user_id`, `product_id`, `qty`) VALUES (?, ?, ?)"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	_, err = stnt.Exec(user_id, product_id, qty)
	if err != nil {
		return err
	}

	return nil
}

func FetchCartByUserId(user_id int) (Response, error) {
	var obj Cart
	var res Response
	var arrobj []Cart

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `carts` WHERE `user_id` = ?"

	rows, err := con.Query(sqlStatement, user_id)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Cart_id, &obj.User_id, &obj.Product_id, &obj.Qty)

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

func DeleteCart(cart_id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM `carts` WHERE `cart_id` = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(cart_id)
	if err != nil {
		return res, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsDeleted,
	}

	return res, nil
}

func EditQty(cart_id int, new_qty int) (Response, error) {
	var res Response
	con := db.CreateCon()

	// check if cart item exists for specified user_id and cart_id
	var count int
	err := con.QueryRow("SELECT COUNT(*) FROM `carts` WHERE `cart_id` = ?", cart_id).Scan(&count)
	if err != nil {
		return res, err
	}
	if count == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Product not found"
		return res, nil
	}

	// update cart item quantity
	sqlStatement := "UPDATE `carts` SET `qty` = ? WHERE `cart_id` = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(new_qty, cart_id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	if rowsAffected == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Cart item not found"
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

package models

import (
	"database/sql"
	"e-commerce/db"
	"errors"
	"fmt"
	"net/http"
)

type Order struct {
	Order_id   int    `json:"order_id"`
	User_id    int    `json:"user_id"`
	Product_id int    `json:"product_id"`
	Qty        int    `json:"qty"`
	Status     string `json:"status"`
}

func MoveCartToOrder(user_id int) error {
	con := db.CreateCon()

	tx, err := con.Begin()
	if err != nil {
		return err
	}

	// Check if there are items in the user's cart
	sqlStatement := "SELECT COUNT(*) FROM `carts` WHERE `user_id` = ?"
	var count int
	err = con.QueryRow(sqlStatement, user_id).Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}
	if count == 0 {
		tx.Rollback()
		return fmt.Errorf("no items found in cart")
	}

	// Get all items in the user's cart
	sqlStatement = "SELECT `product_id`, `qty` FROM `carts` WHERE `user_id` = ?"
	rows, err := con.Query(sqlStatement, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Prepare statement to insert items into orders table
	insertStmt, err := tx.Prepare("INSERT INTO `orders` (`user_id`, `product_id`, `qty`, `status`) VALUES (?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	// Prepare statement to delete items from carts table
	deleteStmt, err := tx.Prepare("DELETE FROM `carts` WHERE `user_id` = ? AND `product_id` = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	// Move items from cart to order
	for rows.Next() {
		var product_id int
		var qty int
		err = rows.Scan(&product_id, &qty)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = insertStmt.Exec(user_id, product_id, qty, "pending")
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = deleteStmt.Exec(user_id, product_id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = rows.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = insertStmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = deleteStmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func FetchOrderByUserId(user_id int) (Response, error) {
	var obj Order
	var res Response
	var arrobj []Order

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `orders` WHERE `user_id` = ?"

	rows, err := con.Query(sqlStatement, user_id)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Order_id, &obj.User_id, &obj.Product_id, &obj.Qty, &obj.Status)

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

func DeleteOrder(user_id int, order_id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM `orders` WHERE `user_id` = ? AND `order_id` = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(user_id, order_id)
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

func UpdateOrderStatus(order_id int, user_id int, product_id int, status string) (Response, error) {
	con := db.CreateCon()
	var res Response

	sqlStatement := "UPDATE `orders` SET `status` = ? WHERE `user_id` = ? AND `order_id` = ?"

	stnt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stnt.Exec(status, user_id, order_id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func GetProductIdByOrderId(order_id int) (int, error) {
	con := db.CreateCon()

	sqlStatement := "SELECT `product_id` FROM `orders` WHERE `order_id` = ?"
	row := con.QueryRow(sqlStatement, order_id)

	var product_id int
	err := row.Scan(&product_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("order not found")
		}
		return 0, err
	}

	return product_id, nil
}

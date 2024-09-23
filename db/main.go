package db

import (
	"clean-architecture/db/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", "host=db user=postgres password=postgres dbname=orders port=5432 sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Database connected")
	}
	return db
}

func ListOrders() ([]models.Order, error) {
	var orders []models.Order
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Products, &order.TotalAmount); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

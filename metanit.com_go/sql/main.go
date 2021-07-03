package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into productdb.Products (model, company, price) vaules (?, ?, ?)", "iPhone", "Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

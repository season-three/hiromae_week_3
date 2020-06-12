package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var id int
var name string
var price int
var err error
var db *sql.DB

//Fruits 構造体
type Fruits struct {
	id    int
	name  string
	price int
}

func main() {
	fruits := Fruits{}

	//接続
	db, err = sql.Open("mysql", "root:11194222@/gomysql")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//複数レコードをselect -> ループ処理
	rows, err := db.Query("select id, name, price from gomysql")
	if err != nil {
		fmt.Println(err)
	}

	//ループ処理 + Next関数
	for rows.Next() {
		err = rows.Scan(&fruits.id, &fruits.name, &fruits.price)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fruits)

	}

}

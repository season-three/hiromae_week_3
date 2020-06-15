package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

//ID 宣言
var ID int

//Name 宣言
var Name string

//Price 宣言
var Price int
var err error
var db *sql.DB

//Fruits 構造体
type Fruits struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	//Echo立ち上げる
	e := echo.New()
	//GETリクエスト（ルーティング）
	e.GET("/show", show)
	//サーバー起動
	e.Start(":9000")
}

func show(c echo.Context) error {
	//構造体（Fruitsの中にあるフィールドをfruits変数として定義？）
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
	for i := 0; i < 3; i++ {

		rows.Next()
		err = rows.Scan(&fruits.ID, &fruits.Name, &fruits.Price)
		if err != nil {
			fmt.Println(err)
		} else if i == 3 {
			break
		}

	}
	return c.JSON(http.StatusOK, fruits)

	result := run()
	c.JSON(http.StatusOK, result)
}

func run() []Fruits {
	fruits := Fruits{}
	xs := make([]Fruits, 0)
	for _, n := range fruits {
		x := Fruits{ID: n, Name: n, Price: n}
		xs = append(xs, x)
	}
	return xs
}

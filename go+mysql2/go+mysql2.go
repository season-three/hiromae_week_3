package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

//Fruit 構造体
type Fruit struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

//ID 宣言
var ID int

//Name 宣言
var Name string

//Price 宣言
var Price int
var err error
var db *sql.DB
var fruits string

func main() {
	//Echo立ち上げる
	e := echo.New()
	//GETリクエスト（ルーティング）
	e.GET("/show", show)
	//サーバー起動
	e.Start(":9000")
}
func show(c echo.Context) error {

	fruits := []*Fruit{
		{
			ID:    ID,
			Name:  Name,
			Price: Price,
		},
	}

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
		err = rows.Scan(&ID, &Name, &Price)
		if err != nil {
			fmt.Println(err)
			if err != nil {
				fmt.Println(err)
			} else if i == 3 {
				break
			}
		}
		fruits = append(fruits)

	}
	return c.JSON(http.StatusOK, fruits)

}

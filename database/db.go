package main

import (
	"database/sql"
	"fmt"
    "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func connect() (*sql.DB, error) {
    // [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
    dbconf := "todo:password@tcp(127.0.0.1:3306)/todo?charset=utf8mb4"

    db, err := sql.Open("mysql", dbconf)

    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}

func saveData(c echo.Context) error {
    text := c.FormValue("text")

    db,err := connect()
    if err != nil {
        fmt.Println(err)
        return c.String(http.StatusInternalServerError, "保存までにエラーが発生しました")
    }
    defer db.Close()

    query := "INSERT INTO users (name, age) VALUES (?, ?)"
    if _ ,err := db.Exec(query, text); err != nil {
        fmt.Println(err)
        return c.String(http.StatusInternalServerError, "保存までにエラーが発生しました")
    }

    return c.String(http.StatusOK, "保存完了")
}
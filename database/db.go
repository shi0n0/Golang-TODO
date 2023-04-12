package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
    dbconf := "todo:password@tcp(127.0.0.1:3306)/todo?charset=utf8mb4"

    db, err := sql.Open("mysql", dbconf)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    defer func() {
        if err := db.Close(); err != nil {
            fmt.Println(err.Error())
        }
    }()

    err = db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    }

    fmt.Println("データベース接続成功")
}

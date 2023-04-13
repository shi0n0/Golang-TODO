package database

import (
	"database/sql"
	"fmt"
    "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func CreateTable() error {
    db, err := Connect()
    if err != nil {
        return err
    }
    defer db.Close()

    // テーブル作成のクエリー
    query := `CREATE TABLE IF NOT EXISTS your_table_name (
                    id INT NOT NULL AUTO_INCREMENT,
                    content VARCHAR(255) NOT NULL,
                    PRIMARY KEY (id)
                ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

    // クエリーを実行
    _, err = db.Exec(query)
    if err != nil {
        return err
    }

    fmt.Println("テーブルが正常に作られました")
    return nil
}

func Connect() (*sql.DB, error) {
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

func SaveData(c echo.Context) error {
    content := c.FormValue("text")

    db,err := Connect()
    if err != nil {
        fmt.Println(err)
        return c.String(http.StatusInternalServerError, "保存までにエラーが発生しました")
    }
    defer db.Close()

    query := "INSERT INTO users (text) VALUES (?)"
    if _ ,err := db.Exec(query, content); err != nil {
        fmt.Println(err)
        return c.String(http.StatusInternalServerError, "保存までにエラーが発生しました")
    }

    return c.String(http.StatusOK, "保存完了")
}
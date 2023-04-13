package main

import (
	"fmt"
	"log"

	"github.com/shi0n0/Golang-TODO/database"
)

func main() {
	if err := CreateTable(); err != nil {
		log.Fatalf("作成に失敗しました: %v", err)
	}
}

func CreateTable() error {
    db, err := database.Connect()
    if err != nil {
        return err
    }
    defer db.Close()

    // テーブル作成のクエリー
    query := `CREATE TABLE IF NOT EXISTS content (
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
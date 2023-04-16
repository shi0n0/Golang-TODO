package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shi0n0/Golang-TODO/database"
	"github.com/shi0n0/Golang-TODO/handler"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// レンダリングエンジンの設定
	render := &TemplateRender{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = render

	// ルーティングの設定
	e.GET("/", handler.ShowHTML)
	e.GET("/save", func(c echo.Context) error {
		// データをMySQLデータベースに保存する
		if err := database.SaveData(c); err != nil {
			return c.String(http.StatusInternalServerError, "データの保存に失敗")
		}
		return c.String(http.StatusOK, "データの保存完了")
	})

	// サーバーの起動
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("サーバー起動に失敗しました: %v", err)
		os.Exit(1)
	}

	// データベースの接続と終了処理
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました: %v", err)
		os.Exit(1)
	}
	defer db.Close()
}

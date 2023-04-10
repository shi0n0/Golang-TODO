package handler

import(
	"net/http"
	"github.com/labstack/echo"
)

func ShowHTML(c echo.Context) error {
	Text := map[string]string{
		"content":"ダミー",
	}
	return c.Render(http.StatusOK, "top", Text)
}
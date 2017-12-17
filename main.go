package main

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	//QTraderURL for qtrader.io home page URL
	QTraderURL = "https://qtrader.io"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Redirect(301, QTraderURL)
		return nil
	})
	e.GET("/:pkg", func(c echo.Context) error {
		// Package name from path `/:pkg`
		pkg := c.Param("pkg")
		isGoGet := c.QueryParam("go-get")
		if pkg != "" && isGoGet == "1" {
			htmlstring := `<html>
	<head>
        <meta name="go-import" content="qtrx.io/` + pkg + ` git https://github.com/q-trader/` + pkg + `">
        <meta name="go-source" content="qtrx.io/` + pkg + `     https://github.com/q-trader/` + pkg + ` https://github.com/q-trader/` + pkg + `/tree/master{/dir} https://github.com/q-trader/` + pkg + `/blob/master{/dir}/{file}#L{line}">
    </head>
</html>
`
			return c.HTML(http.StatusOK, htmlstring)

		}
		c.Redirect(301, QTraderURL)
		return nil

	})
	e.Logger.Fatal(e.Start(":1323"))
}

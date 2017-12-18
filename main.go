package main

import (
	"net/http"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

const (
	//QTraderURL for qtrader.io home page URL
	QTraderURL = "https://qtrader.io"
)

type PackageView struct {
	Title string
	Name  string
}
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		c.Redirect(301, QTraderURL)
		return nil
	})
	//support qtrx.io/x,qtrx.io/x/y,qtrx.io/x/y/z, qtrx.io/x/y/z/
	e.GET("/:pkg", handPkg)
	e.GET("/:pkg/:subPkg", handPkg)
	e.GET("/:pkg/:subPkg/:sSubPkg", handPkg)
	e.GET("/:pkg/:subPkg/:sSubPkg/:sSSubPkg", handPkg)

	e.Logger.Fatal(e.Start(":1323"))
}

func handPkg(c echo.Context) error {
	pkgName := c.Param("pkg")
	isGoGet := c.QueryParam("go-get")
	pkg := &PackageView{
		Title: pkgName,
		Name:  pkgName,
	}
	if pkg.Name != "" && isGoGet == "1" {
		return c.Render(http.StatusOK, "pkg.html", pkg)

	}
	c.Redirect(301, QTraderURL)
	return nil
}

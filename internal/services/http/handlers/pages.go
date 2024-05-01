package handlers

import (
	"errors"
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yannis94/kanshi/pkg"
)

type PageHandler struct{}
type HTMLTemplate struct {
	templates map[string]*template.Template
}

func NewHTMLTemplate(dirPath string) *HTMLTemplate {
	templates := make(map[string]*template.Template)
	templateFiles, err := pkg.ListFiles(dirPath)

	if err != nil {
		panic(err)
	}

	baseTempl := filepath.Join(dirPath, "layouts", "main.html")
	for _, file := range templateFiles {
		if filepath.Ext(file) != ".html" {
			continue
		}
		var templatePath string
		if s := strings.Split(file, dirPath); len(s) > 0 {
			templatePath = s[1]
		} else {
			templatePath = file
		}
		templates[templatePath] = template.Must(template.ParseFiles(file, baseTempl))
	}

	return &HTMLTemplate{templates: templates}
}

func (t *HTMLTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	templ, ok := t.templates[name]
	if !ok {
		return errors.New("Template not found -> " + name)
	}
	return templ.ExecuteTemplate(w, "main.html", data)
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) Index(c echo.Context) error {
	return c.Render(200, "/views/index.html", map[string]interface{}{})
}

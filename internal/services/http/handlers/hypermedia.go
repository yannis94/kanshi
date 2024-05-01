package handlers

import "html/template"

type HypermediaTemplate struct {
	templates *template.Template
}

func NewHypermediaTemplate(dirPath string) *HypermediaTemplate {
	templates := template.Must(template.ParseGlob(dirPath + "/*.html"))
	return &HypermediaTemplate{templates: templates}
}

func (t *HypermediaTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

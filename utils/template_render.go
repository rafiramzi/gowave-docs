package utils

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewRenderer(path string) *TemplateRenderer {
	templates, err := template.ParseGlob(path)
	if err != nil {
		panic(err)
	}
	return &TemplateRenderer{
		templates: templates,
	}
}



func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Try full name first
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		// Try finding by filename only
		for _, tmpl := range t.templates.Templates() {
			if filepath.Base(tmpl.Name()) == name {
				return tmpl.Execute(w, data)
			}
		}
	}
	return err
}


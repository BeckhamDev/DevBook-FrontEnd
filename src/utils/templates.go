package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplates(){
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
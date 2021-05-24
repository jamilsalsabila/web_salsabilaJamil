package routes

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*"))
}

func () {
	mux := httprouter.New()

	mux.GET("/", indeks)

	http.ListenAndServe(":8080", mux)
}

func indeks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "indeks.gohtml", nil)
}


package functions

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func Render(v any, w http.ResponseWriter, file ...string) {
	t, err := template.ParseFiles("assets/html/" + filepath.Join(file...))
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, v)
	if err != nil {
		log.Println(err)
		return
	}
}

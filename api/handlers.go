package api

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}



func addLink(w http.ResponseWriter, r *http.Request) {
  
}

func showLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("showing ...")
}
func redirect(w http.ResponseWriter, r *http.Request) {

}

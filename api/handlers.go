package api

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
  // TODO check if there is a user session , print welcom $name user
	//u := getUser(w, req)
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Alert")
	if err != nil {
		log.Fatalln(err)
	}
}

func singUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		fmt.Println(e, p)
		// TODO check if user already exists
		// TODO check if there is a session for this user
		// creating and setting session
		sID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		// TODO add session to db
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}
func logout(w http.ResponseWriter, r *http.Request) {

}
func addLink(w http.ResponseWriter, r *http.Request) {

}
func showLinks(w http.ResponseWriter, r *http.Request) {
  fmt.Println("showing ...")
}
func redirect(w http.ResponseWriter, r *http.Request) {

}

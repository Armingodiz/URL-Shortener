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
	cookie, err := r.Cookie("session")
	var err2 error
	if err != nil { // check if we have a session cookie
		err2 = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	} else {
		userName := shortener.db.GetSessionInfo(cookie.Value) // user name which blongs to this session cookie
		err2 = tpl.ExecuteTemplate(w, "index.gohtml", userName)
	}
	if err2 != nil {
		log.Fatalln(err2)
	}
}

func singUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		err := shortener.db.SignUp(e, p)
		if err != nil {
      // if user is already exists or there is a session for this user
      fmt.Println(err)
		} else {
			// creating and setting session
			sID := uuid.NewV4()
			cookie := &http.Cookie{
				Name:  "session",
				Value: sID.String(),
			}
			http.SetCookie(w, cookie)
			// adding session to db
			shortener.db.AddSession(cookie.Value, e)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {
	// checking if we have a session
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatalln(err)
	}
	// removing session from db
	shortener.db.Logout(cookie.Value)
	cook := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cook)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addLink(w http.ResponseWriter, r *http.Request) {

}
func showLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("showing ...")
}
func redirect(w http.ResponseWriter, r *http.Request) {

}

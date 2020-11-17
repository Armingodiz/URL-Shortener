package api

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"tawesoft.co.uk/go/dialog"
)


/////////////////////////////////// loading main page :

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

//////////////////////////////// singning up user :

func singUp(w http.ResponseWriter, r *http.Request) {
	// checking if we have a session
	_, err := r.Cookie("session")
	if err == nil {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		err := shortener.db.SignUp(e, p)
		if err != nil {
			// if user is already exists or there is a session for this user
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

////////////////////////////////////// loging in user :

func login(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err == nil {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		e := r.FormValue("email")
		p := r.FormValue("password")
		err = shortener.db.Login(e, p)
		if err != nil {
			log.Fatalln(err)
		}
		// if everything was fine :
		// creates session , set it to cookie and then add session to database
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		shortener.db.AddSession(c.Value, e)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

///////////////////////////////////////////////// logging out user :

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

package api

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"tawesoft.co.uk/go/dialog"
)

/*func loginPage(w http.ResponseWriter, req *http.Request) {
	if checkLog(req) {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var loggedUser user
	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		p := req.FormValue("pass")
		loggedUser, ok := users[e]
		if !ok {
			http.Error(w, "WRONG EMAIL ! ", http.StatusForbidden)
			//dialog.Alert("THERE IS ALREADY AN ACCOUNT WITH THIS EMAIL ! ")
			return
		}
		if loggedUser.password != p {
			http.Error(w, "WRONG password ! ", http.StatusForbidden)
			//dialog.Alert("THERE IS ALREADY AN ACCOUNT WITH THIS EMAIL ! ")
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		sessions[c.Value] = session{e}
		http.Redirect(w, req, "/main", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "logPage.gohtml", loggedUser)
	errHandler(err)
}

func signPage(w http.ResponseWriter, req *http.Request) {
	if checkLog(req) {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var newUser user
	if req.Method == http.MethodPost {
		n := req.FormValue("name")
		e := req.FormValue("email")
		p := req.FormValue("password")
		fmt.Println(n, e, p)
		if _, ok := users[e]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			//dialog.Alert("THERE IS ALREADY AN ACCOUNT WITH THIS EMAIL ! ")
			return
		}
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		sessions[cookie.Value] = session{e}
		newUser = user{n, e, p}
		users[e] = newUser
		db, err := sql.Open("mysql",
			"root:armin3011@tcp(127.0.0.1:3306)/loginSystem")
		errHandler(err)
		stmt, err := db.Prepare(`INSERT INTO users VALUES (?,?,?,?);`)
		defer stmt.Close()
		defer db.Close()
		r, err := stmt.Exec(10, n, e, p)
		errHandler(err)
		ro, err := r.RowsAffected()
		errHandler(err)
		fmt.Println("INSERTED RECORD", ro)
		errHandler(err)
		http.Redirect(w, req, "/main", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "sign.gohtml", newUser)
	errHandler(err)
}
func logOut(w http.ResponseWriter, req *http.Request) {
	if !checkLog(req) {
		dialog.Alert("YOU ARE NOT LOGGED IN YET ! ")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	c, err := req.Cookie("session")
	errHandler(err)
	delete(sessions, c.Value)
	cook := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cook)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}*/
/////////////////////////////////////////////////////////////////////////////////////////////////////////////// main part :

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

////////////////////////////////////// loging in user :

func login(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err == nil {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		e := r.FormValue("email")
		p := r.FormValue("pass")
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

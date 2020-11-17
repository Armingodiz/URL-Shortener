package api

import (
	//"fmt"
	uuid "github.com/satori/go.uuid"
	//"html/template"
	"net/http"
	//	"tawesoft.co.uk/go/dialog"
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

func getUser(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		sID:= uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, cookie)
}

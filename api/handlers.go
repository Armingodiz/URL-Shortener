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
	var err error
	url := r.FormValue("url")
	// remove www.
	url = strings.Replace(url, "www.", "", 1)
	// adding http to url if it doesn't contains it
	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}
	// create shortenned version
	link := shortenUrl(url)
	cookie, err := r.Cookie("session")
	userName := shortener.db.GetSessionInfo(cookie.Value) // user name which blongs to this session cookie
	go func() {
		shortener.db.AddLink(url, link, userName)
	}()
	err := tpl.ExecuteTemplate(w, "result.gohtml", link)
	if err != nil {
		log.Fatalln(err)
	}
}

func showLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("showing ...")
}
func redirect(w http.ResponseWriter, r *http.Request) {

}

func shortenUrl(url string) string {

}

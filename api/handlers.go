package api

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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
	link := r.Host + "/$/" + shortenUrl(url)
	fmt.Println(r.Host)
	cookie, err := r.Cookie("session")
	userName := shortener.db.GetSessionInfo(cookie.Value) // user name which blongs to this session cookie
	go func() {
		shortener.db.AddLink(url, link, userName)
	}()
	err = tpl.ExecuteTemplate(w, "result.gohtml", link)
	if err != nil {
		log.Fatalln(err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	url := r.Host + r.URL.String()
	originalLink := shortener.db.GetLink(url)
	if originalLink != "" {
		http.Redirect(w, r, originalLink, http.StatusSeeOther)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func shortenUrl(url string) string {
	md5 := md5.Sum([]byte(url))
	hash := strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(md5[:])[:6], "/", "_"), "+", "-")
	return hash
}

func showLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("showing ...")
}

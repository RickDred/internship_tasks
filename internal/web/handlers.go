package web

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type errorInfo struct {
	Status int
	Text   string
}

// this for ui, cause i don't want to start other server to implement UI

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/templates/login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/templates/register.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func ErrorHandler(w http.ResponseWriter, statusCode int) {
	tmpl, err := template.ParseFiles("./ui/templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	statusText := http.StatusText(statusCode)
	if err := tmpl.Execute(w, errorInfo{statusCode, statusText}); err != nil {
		log.Println(err)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/templates/profile.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func StaticFilesHandler(w http.ResponseWriter, r *http.Request) {
	path := "./ui" + r.URL.Path

	fileInfo, err := os.Stat(path)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if fileInfo.IsDir() {
		ErrorHandler(w, http.StatusForbidden)
		return
	}
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/static/")
	http.FileServer(http.Dir("./ui/static/")).ServeHTTP(w, r)
}

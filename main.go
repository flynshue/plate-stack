package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var internalServerErr = "Internal Server Error. Please try again!"

func layoutFiles() []string {
	files, err := filepath.Glob("views/bulma-*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func parseForm(req *http.Request, dst interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	return dec.Decode(dst, req.PostForm)
}

func home(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s: %s\n", req.Method, req.URL)
	t, err := template.ParseFiles(layoutFiles()...)
	if err != nil {
		log.Println("Error parsing layout files")
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
	form := newPlateForm()
	if err := t.ExecuteTemplate(w, "bulma", form); err != nil {
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
}

func calc(w http.ResponseWriter, req *http.Request) {
	form := newPlateForm()
	if err := parseForm(req, &form); err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
	}
	form.findCombinations()
	t, err := template.ParseFiles(layoutFiles()...)
	if err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
	if err := t.ExecuteTemplate(w, "bulma", form); err != nil {
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/calc", calc).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}

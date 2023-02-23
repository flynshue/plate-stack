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

type PlateForm struct {
	Barbell float64 `schema:"barbell"`
	Plate   []Plate `schema:"plate"`
	Total   float64
}

type Plate struct {
	Weight   float64 `schema:"weight"`
	Quantity float64 `schema:"quantity"`
}

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
	if err := t.ExecuteTemplate(w, "bulma", nil); err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
}

func calc(w http.ResponseWriter, req *http.Request) {
	var plateForm PlateForm
	if err := parseForm(req, &plateForm); err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
	}
	t, err := template.ParseFiles(layoutFiles()...)
	if err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
	var total float64
	for _, plate := range plateForm.Plate {
		plate.Quantity *= 2
		total += plate.Weight * plate.Quantity
	}
	plateForm.Total = total + plateForm.Barbell
	if err := t.ExecuteTemplate(w, "bulma", plateForm); err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
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

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

var internalServerErr = "Internal Server Error. Please try again!"

var tmplFuncs = template.FuncMap{
	"add1": func(i int) int { return i + 1 },
}

const sessionName = "plate-stack"
const sessionPlatesKey = "plates"

var store = sessions.NewCookieStore([]byte(sessionKey()))

func sessionKey() string {
	if k := os.Getenv("SESSION_KEY"); k != "" {
		return k
	}
	return "plate-stack-default-key-change-me"
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

func loadSessionQuantities(req *http.Request) map[string]int {
	session, err := store.Get(req, sessionName)
	if err != nil {
		return nil
	}
	data, ok := session.Values[sessionPlatesKey].(string)
	if !ok {
		return nil
	}
	var quantities map[string]int
	if err := json.Unmarshal([]byte(data), &quantities); err != nil {
		return nil
	}
	return quantities
}

func applySessionQuantities(form *PlateForm, quantities map[string]int) {
	for i, p := range form.Plates {
		key := fmt.Sprintf("%g", p.Weight)
		if q, ok := quantities[key]; ok {
			form.Plates[i].Quantity = q
		}
	}
}

func saveSessionQuantities(w http.ResponseWriter, req *http.Request, form *PlateForm) {
	session, err := store.Get(req, sessionName)
	if err != nil {
		return
	}
	quantities := make(map[string]int, len(form.Plates))
	for _, p := range form.Plates {
		quantities[fmt.Sprintf("%g", p.Weight)] = p.Quantity
	}
	data, err := json.Marshal(quantities)
	if err != nil {
		return
	}
	session.Values[sessionPlatesKey] = string(data)
	session.Save(req, w)
}

func home(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s: %s\n", req.Method, req.URL)
	t, err := template.New("").Funcs(tmplFuncs).ParseFiles(layoutFiles()...)
	if err != nil {
		log.Println("Error parsing layout files")
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
	form := newPlateForm()
	if quantities := loadSessionQuantities(req); quantities != nil {
		applySessionQuantities(form, quantities)
	}
	if err := t.ExecuteTemplate(w, "bulma", form); err != nil {
		log.Println(err)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}
}

func calc(w http.ResponseWriter, req *http.Request) {
	form := newPlateForm()
	if err := parseForm(req, form); err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
	}
	saveSessionQuantities(w, req, form)
	form.findCombinations()
	t, err := template.New("").Funcs(tmplFuncs).ParseFiles(layoutFiles()...)
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
	r.HandleFunc("/", calc).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}

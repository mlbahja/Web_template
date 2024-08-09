package main

import (
	"fmt"
	"html/template"
	"net/http"

	utils "ascii_web/utils"
)

func AsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405: Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == "POST" {
		data := r.PostFormValue("textInput")
		banner := r.PostFormValue("bannerType")
		if len(data) == 0 || len(banner) == 0 {
			http.Error(w, "400: Bad request.", http.StatusBadRequest)
			return
		}
		result, check := utils.AsciiArtGenerator(data, banner)
		if check == 1 {
			http.Error(w, "400: Bad request.", http.StatusBadRequest)
			return
		}
		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, result)
		if err != nil {
			http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
			return
		}
	}
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404: Page not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405: Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/ascii-art", AsciiArtResult)
	fmt.Println("\033[32mServer started at http://127.0.0.1:8080\033[0m")
	http.ListenAndServe(":8080", nil)
}

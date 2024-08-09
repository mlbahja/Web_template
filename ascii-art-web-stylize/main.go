package main

import (
	"fmt"
	"html/template"
	"net/http"

	// "strings"

	utils "ascii_web/utils"
)

type errorType struct {
	ErrorCode string
	Message   string
}

func AsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorPages(w, 405)
		return
	}
	if r.Method == "POST" {
		data := r.PostFormValue("textInput")
		banner := r.PostFormValue("bannerType")
		if len(data) == 0 || len(banner) == 0 {
			errorPages(w, 400)
			return
		}
		result, check := utils.AsciiArtGenerator(data, banner)
		if check == 1 {
			errorPages(w, 400)
			return
		}
		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			errorPages(w, 500)
			return
		}
		err = t.Execute(w, result)
		if err != nil {
			errorPages(w, 500)
			return
		}
	}
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPages(w, 405)
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errorPages(w, 500)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		errorPages(w, 500)
		return
	}
}

func errorPages(w http.ResponseWriter, code int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		return
	} else if code == 404 {
		w.WriteHeader(http.StatusNotFound)
		err = t.Execute(w, errorType{ErrorCode: "404", Message: "Sorry, the page you are looking for does not exist."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 405 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err = t.Execute(w, errorType{ErrorCode: "405", Message: "Method not allowed."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
	}
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style/" {
		errorPages(w, 404)
		return
	}
	fs := http.FileServer(http.Dir("./style"))
	http.StripPrefix("/style/", fs).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/style/", serveCSS)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorPages(w, 404)
			return
		}
		RootPage(w, r)
	})

	http.HandleFunc("/ascii-art", AsciiArtResult)
	fmt.Println("\033[32mServer started at http://127.0.0.1:8080\033[0m")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func AsciiArtGenerator(text string) string {
	// Simple ASCII art generator (mock)
	asciiArt := strings.ToUpper(text)
	return asciiArt
}

func AsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405: Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	data := r.PostFormValue("textInput")
	if len(data) == 0 {
		http.Error(w, "400: Bad request.", http.StatusBadRequest)
		return
	}
	result := AsciiArtGenerator(data)

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

func exportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405: Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	asciiArt := r.FormValue("ascii-art")
	if asciiArt == "" {
		http.Error(w, "No ASCII art provided", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(asciiArt)))
	w.Write([]byte(asciiArt))
}

func main() {
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/ascii-art", AsciiArtResult)
	http.HandleFunc("/export", exportHandler)
	fmt.Println("Server started at http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

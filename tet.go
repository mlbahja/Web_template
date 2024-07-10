package main

import (
	"net/http"
	"strconv"
	"text/template"
)

func pageweb(w http.ResponseWriter, r *http.Request) {
	templatepage, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "400: Bad request.", http.StatusBadRequest)
		return
	}
	nbr1, _ := strconv.Atoi(r.FormValue("num1"))
	nbr2, _ := strconv.Atoi(r.FormValue("num2"))
	result := nbr1 + nbr2
	templatepage.Execute(w, result)
}

func main() {
	http.HandleFunc("/", pageweb)
	http.ListenAndServe(":8080", nil)
}

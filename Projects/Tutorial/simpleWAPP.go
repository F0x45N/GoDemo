package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, My first Go Web!\n")
	fmt.Fprintf(w, "This is Neat!")
}

func cred_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Credit to Sentdex for this demo")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/credit", cred_handler)
	http.ListenAndServe(":8000", nil)
}
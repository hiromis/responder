package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go to /code/{code} to receive a response with the status code")
}

func code(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, ok := vars["code"]
	if !ok {
		fmt.Println("code is missing in parameters")
	}
	code, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println("code in parameters must be numeric")
	}

	text := http.StatusText(code)
	if text == "" {
		text = "Unknown code"
	}

	w.WriteHeader(code)
	w.Write([]byte(text))
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/code/{code}", code)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}

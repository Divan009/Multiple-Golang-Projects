package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successfully")
	name := r.FormValue("name")
	address := r.FormValue("address")
	_, err := fmt.Fprintf(w, "Name = %s\n", name)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		//http.NotFound(w, r)
		http.Error(w, "This page is not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		//http.StatusMethodNotAllowed()
		http.Error(w, "Method is not Supported", http.StatusMethodNotAllowed)
		return
	}

	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(
		http.Dir("./static"),
	)

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

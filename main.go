package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parse form error : %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	Name := r.FormValue("name")
	phone_Number := r.FormValue("phone")

	fmt.Fprintf(w, "Name = %s\n", Name)
	fmt.Fprintf(w, "Phone Number= %s\n", phone_Number)

}

// wr is for requets and response from the server and the user
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// to double check whether the request is made through /hello or not
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return

	}

	// to check only GET Method works on /hello
	if r.Method != "GET" {
		http.Error(w, "Method not found ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello User")

}
func main() {

	fileServer := http.FileServer(http.Dir("./static")) // this will look for index.html file

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at 8080 port\n")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}

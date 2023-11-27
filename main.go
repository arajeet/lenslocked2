package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Header)
	fmt.Fprint(w, "<h1> Welcome to my website <h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

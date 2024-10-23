package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye World")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)

	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

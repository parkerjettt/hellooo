package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", server)
	fmt.Println("Server is running!")
	http.ListenAndServe(":1234", nil)
}

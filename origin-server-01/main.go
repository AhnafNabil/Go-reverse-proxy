package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from origin Server 1")
	})

	fmt.Println("Origin Server 1 is listening on port 8081...")
	http.ListenAndServe(":8081", nil)
}

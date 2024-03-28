package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from origin Server 2")
	})

	fmt.Println("Origin Server 2 is listening on port 8082...")
	http.ListenAndServe(":8082", nil)
}

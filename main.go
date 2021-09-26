package main

import (
	"net/http"
		"os"	
	"fmt"	
)

func main() {
		port := os.Getenv("PORT")
	http.HandleFunc("/", foo)
	http.ListenAndServe(":"+port, nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Status code", "200")
	w.WriteHeader(200)
	fmt.Fprintf("test")
}

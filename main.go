package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/api", api)
	http.ListenAndServe(":"+port, nil)
}

func api(w http.ResponseWriter, r *http.Request) {
	log.Print("new request")
	query := r.URL.Query()

	token_byte := query.Get("3c6e0b8a9c15224a8228b9a98ca1531d")

	token := string(token_byte[:])
	if token == "ac01a9a846016b13e1249040c3bb1c3e" {

		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(200)
	}

}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Cardinal struct {
	Status string
	Code   []string
}

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
		mss := Cardinal{"Success", []string{"201"}}

		js, err := json.Marshal(mss)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		log.Print("Success: 201")
	} else {
		mss := Cardinal{"ErrAuth", []string{"401"}}
		js, err := json.Marshal(mss)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		log.Print("ErrAuth: 401")
	}

}

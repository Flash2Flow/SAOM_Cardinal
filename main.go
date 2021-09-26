package main

import (
	"net/http"
		"os"	
		"encoding/json"
)

type Profile struct {
	Name		string
	Hobbies []string
}

func main() {
		port := os.Getenv("PORT")
	http.HandleFunc("/", foo)
	http.ListenAndServe(":"+port, nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

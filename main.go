package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gopkg.in/gomail.v2"
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

	token_byte := query.Get("token")
	email_byte := query.Get("email")
	pass_byte := query.Get("new_pass")

	token := string(token_byte[:])
	email := string(email_byte[:])
	pass := string(pass_byte[:])
	if token == "ac01a9a846016b13e1249040c3bb1c3e" {

		m := gomail.NewMessage()
		m.SetHeader("From", "tmushkaterova@gmail.com")
		m.SetHeader("To", email)
		m.SetHeader("Subject", "SAOM ONLINE Drop Password!")
				m.SetBody("text/html", "Ваш новый пароль, изменить его вы можете в личном кабинете ( в разработке ) - ("+pass+")")

		d := gomail.NewPlainDialer("smtp.gmail.com", 587, "tmushkaterova@gmail.com", "537003DOsaV")
		if err := d.DialAndSend(m); err != nil {
			log.Print(err)
			mss := Cardinal{"Internal Server Error", []string{"500"}}
			js, err := json.Marshal(mss)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("MSS", "Cardinal")
			w.Write(js)
			log.Print("Internal Server Error: 500")
		}

		mss := Cardinal{"Success", []string{"201"}}

		js, err := json.Marshal(mss)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("MSS", "Cardinal")
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
		w.Header().Set("MSS", "Cardinal")
		w.Write(js)
		log.Print("ErrAuth: 401")
	}

}


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
	access_drop_byte := query.Get("access_drop")

	token := string(token_byte[:])
	email := string(email_byte[:])
	access_drop := string(access_drop_byte[:])
	if token == "ac01a9a846016b13e1249040c3bb1c3e" {

		m := gomail.NewMessage()
		m.SetHeader("From", "tmushkaterova@gmail.com")
		m.SetHeader("To", email)
		m.SetHeader("Subject", "SAOM ONLINE Drop Password!")
		m.SetBody("text/html", "Перейдите по ссылке что бы сбросить пароль - <a href='http://ch37276.tmweb.ru/drop.php?email="+email+"&access_drop="+access_drop+">Link</a>")

		d := gomail.NewPlainDialer("smtp.gmail.com", 587, "tmushkaterova@gmail.com", "537003DOsaV")
		if err := d.DialAndSend(m); err != nil {
			panic(err)
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

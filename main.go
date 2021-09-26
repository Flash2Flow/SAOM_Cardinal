package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
)

type Cardinal struct {
	Status string
	Code   []string
}

type Drop struct {
	Access_drop string
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

	token := string(token_byte[:])
	email := string(email_byte[:])
	if token == "ac01a9a846016b13e1249040c3bb1c3e" {

		db, err := sql.Open("mysql", "re-incarnation:537003ReIncar$Mailen@tcp(37.146.123.137:3306)/poplavok")

		if err != nil {
			log.Print(err)
		}

		defer db.Close()

		res, err := db.Query("SELECT `access_drop` FROM `users` WHERE `email`=?", email)
		if err != nil {
			log.Print(err)
		}

		drop_full := []Drop{}
		for res.Next() {
			var drop Drop
			err = res.Scan(&drop.Access_drop)
			if err != nil {
				log.Print(err)
			}
			drop_full = append(drop_full, drop)
		}

		for _, drop := range drop_full {
			var Body_url = "Перейдите по ссылке что бы сбросить пароль - <a href='http://ch37276.tmweb.ru/drop.php?email=" + email + "&access_drop=" + drop.Access_drop + "'>Link</a>"
			m := gomail.NewMessage()
			m.SetHeader("From", "tmushkaterova@gmail.com")
			m.SetHeader("To", email)
			m.SetHeader("Subject", "SAOM ONLINE Drop Password!")
			m.SetBody("text/html", Body_url)

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
		}
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

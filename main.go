package main

import (
	"log"
	"net/http"
	"net/smtp"
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
	email_byte := query.Get("0c83f57c786a0b4a39efab23731c7ebc")
	token_byte := query.Get("3c6e0b8a9c15224a8228b9a98ca1531d")
	access_drop_byte := query.Get("c1a8a39a96d32cac85fd7bca0d50830b")

	email := string(email_byte[:])
	access_drop := string(access_drop_byte[:])
	token := string(token_byte[:])

	if token == "ac01a9a846016b13e1249040c3bb1c3e" {
		send("Перейдите по ссылке что бы сбросить пароль - http//site.com/drop.php?0c83f57c786a0b4a39efab23731c7ebc=%s&c1a8a39a96d32cac85fd7bca0d50830b=%s", email, access_drop)
	}

}
func send(body string, email string, access_drop string) {
	from := "tmushkaterova@gmail.com"
	pass := "537003DOSAV"
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: SAOM ONLINE Drop Password\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Println("StatusRequest: 401")
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("StatusRequest: 201")
	log.Println("Sent to %s, ", to)
}

package main

import (
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth("", "sebastian.1701811396@ucaldas.edu.co", "ecwd lelv ritz lmgw", "smtp.gmail.com")

	to := []string{"sebastian.1701811396@ucaldas.edu.co"}

	msg := []byte(
		"To: sebastian.1701811396@ucaldas.edu.co\r\n" + 
		"Subject: Hello from Golang\r\n" +
		"\r\n" +
		"Hola mundo")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "sebastian.1701811396@ucaldas.edu.co", to, msg);

	if err != nil {
		log.Fatal(err)
	}
}

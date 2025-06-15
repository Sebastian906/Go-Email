package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"sebastian.1701811396@ucaldas.edu.co",
		"ecwd lelv ritz lmgw",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"sebastian.1701811396@ucaldas.edu.co",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendMailHtml(subject string, templatePath string, to []string) {

	// Obtener Html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct{Name string}{Name: "Sebastián"})

	auth := smtp.PlainAuth(
		"",
		"sebastian.1701811396@ucaldas.edu.co",
		"ecwd lelv ritz lmgw",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"sebastian.1701811396@ucaldas.edu.co",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// sendMailSimple(
	// 	"Asunto", 
	// 	"Cuerpo de texto", 
	// 	[]string{"sebastian.1701811396@ucaldas.edu.co"},
	// )

	sendMailHtml(
		"Asunto",
		"./test.html",
		[]string{"sebastian.1701811396@ucaldas.edu.co"},
	)
}
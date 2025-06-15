package main

import "gopkg.in/gomail.v2"

func main() {

	m := gomail.NewMessage()

	m.SetHeader("From", "sebastian.1701811396@ucaldas.edu.co")
	m.SetHeader("To", "sebastian.1701811396@ucaldas.edu.co", "elsebas1912@gmail.com")
	m.SetAddressHeader("Cc", "elsebas1912@gmail.com", "Sebastián")
	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", "Hello <b>Sebastián</b> <br> Chekout this little dragon")
	m.Attach("C:/Users/usuario/Downloads/1749740384614.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "sebastian.1701811396@ucaldas.edu.co", "ecwd lelv ritz lmgw")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
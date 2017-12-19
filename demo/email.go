package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "810169879@qq.com", "wangzhen@topsec.com.cn")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("192.168.74.128", 25, "user", "123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

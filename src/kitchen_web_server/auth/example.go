package auth

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

var (
	from string
	pass string
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func SendEmail(to, newusername, newpassword string) {
	// subject, message string
	subject := "Pause Kitchen Management App superuser username and password Change Notification"
	message := "New username is " + newusername + " and password is " + newpassword + ".\n" +
		"If you want to stop these notifications, please remove your username at the super admins editing portion of the App."

	// ADD VALID GMAIL CREDENTIALS
	// You will also need to disable secure login on Google for this to work.
	from = "stolafbigdisk@gmail.com"
	pass = "sk3q&sxfV&K7{^"
	toarr := []string{}
	toarr = append(toarr, to)
	mail := Mail{from, toarr, subject, message}
	messageBody := mail.BuildMessage()
	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}
	auth := smtp.PlainAuth("", mail.senderId, pass, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Email sent successfully")
}

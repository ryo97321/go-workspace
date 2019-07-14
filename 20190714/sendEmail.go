package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()

	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("From : ")
	stdin.Scan()
	From := stdin.Text()

	fmt.Print("From password : ")
	stdin.Scan()
	password := stdin.Text()

	fmt.Print("To : ")
	stdin.Scan()
	To := stdin.Text()

	Bcc := ""
	fmt.Print("Bcc (if no BCC -> \"no\") : ")
	stdin.Scan()
	if stdin.Text() != "no" {
		Bcc = stdin.Text()
	}

	Cc := ""
	fmt.Print("Cc (if no Cc -> \"no\") : ")
	stdin.Scan()
	if stdin.Text() != "no" {
		Cc = stdin.Text()
	}

	Subject := ""
	fmt.Print("Subject (if no Subject -> \"no\") : ")
	stdin.Scan()
	if stdin.Text() != "no" {
		Subject = stdin.Text()
	}

	Text := ""
	fmt.Print("Text (if no Text -> \"no\") : ")
	stdin.Scan()
	if stdin.Text() != "no" {
		Text = stdin.Text()
	}

	HTML := ""
	fmt.Print("HTML (if no HTML -> \"no\") : ")
	stdin.Scan()
	if stdin.Text() != "no" {
		HTML = stdin.Text()
	}

	e.From = From
	e.To = []string{To}
	if Bcc != "" {
		e.Bcc = []string{Bcc}
	}
	if Cc != "" {
		e.Cc = []string{Cc}
	}
	if Subject != "" {
		e.Subject = Subject
	}
	if Text != "" {
		e.Text = []byte(Text)
	}
	if HTML != "" {
		e.HTML = []byte(HTML)
	}
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", From, password, "smtp.gmail.com"))
}

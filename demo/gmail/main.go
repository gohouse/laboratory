package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	var to = "redis111@linshiyouxiang.net"
	err := NewGmail("abc@gmail.com", "wkkqysvpoonhphmc").
		SetSubject("gmail in golang test").
		SetContent("test content").
		Send(to)
	log.Println(err)
}

type From struct {
	Email    string
	Password string
}
type Content struct {
	To          string
	Subject     string
	ContentType string
	Content     string
}
type Gmail struct {
	from    From
	content Content
	host    string
}

func NewGmail(fromEmail, password string) *Gmail {
	return &Gmail{From{
		Email:    fromEmail,
		Password: password,
	}, Content{
		ContentType: "text/html",
	}, "smtp.gmail.com:587"}
}

func (g *Gmail) SetSubject(text string) *Gmail {
	g.content.Subject = text
	return g
}

func (g *Gmail) SetContent(htmlContent string) *Gmail {
	g.content.Content = htmlContent
	return g
}
func (g *Gmail) Send(toEmail string) error {
	return g.SendMulti([]string{toEmail})
}

func (g *Gmail) buildContent() []byte {
	if g.content.Subject == "" {
		var subjectLen = 16
		var content = []rune(g.content.Content)
		if len(content) > subjectLen {
			g.content.Subject = string(content[:16])
		} else {
			g.content.Subject = g.content.Content
		}
	}
	return []byte(fmt.Sprintf("Subject: %v\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n "+
		"%v", g.content.Subject, g.content.Content))
}
func (g *Gmail) SendMulti(toEmails []string) error {
	var content = g.buildContent()
	auth := smtp.PlainAuth("", g.from.Email, g.from.Password, "smtp.gmail.com")
	return smtp.SendMail(g.host, auth, g.from.Email, toEmails, content)
}

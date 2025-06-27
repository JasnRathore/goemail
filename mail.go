package goemail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
	"strconv"
	"strings"
)

type MailProfile struct {
	Name     string //eg: Marketin Bot
	UserName string //Actual mail Account
	From     string //First Last <test@example.com>
	Host     string //smtp.example.com:25
	Password string
}

type MailAttachment struct {
	FileName string
	Data     []byte
}

func NewProfile(name, userName, from, host, password string) MailProfile {
	return MailProfile{
		Name:     name,
		UserName: userName,
		From:     from,
		Host:     host,
		Password: password,
	}
}

func (mp *MailProfile) SendMail(to, subject, body string, attachments []MailAttachment, useTrackingImage bool) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mp.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, att := range attachments {
		data := att.Data
		m.Attach(att.FileName, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(data)
			return err
		}))
	}

	host, port, err := mp.parseHostPort()

	if err != nil {
		return err
	}

	d := gomail.NewDialer(host, port, mp.UserName, mp.Password)
	return d.DialAndSend(m)
}

func (mp *MailProfile) SendTestMail(to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mp.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "This is  Test Email from Go-Email")
	m.SetBody("text/html", "its ok")

	host, port, err := mp.parseHostPort()

	if err != nil {
		return err
	}

	d := gomail.NewDialer(host, port, mp.UserName, mp.Password)
	return d.DialAndSend(m)
}

func (mp *MailProfile) parseHostPort() (string, int, error) {
	parts := strings.Split(mp.Host, ":")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid host:port format")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}
	return parts[0], port, nil
}

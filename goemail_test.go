package goemail

import "testing"

func TestParseHostPort(t *testing.T) {
	mp := MailProfile{Host: "smtp.example.com:587"}
	host, port, err := mp.parseHostPort()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if host != "smtp.example.com" {
		t.Errorf("expected host 'smtp.example.com', got '%s'", host)
	}
	if port != 587 {
		t.Errorf("expected port 587, got %d", port)
	}
}

func TestSendMail_InvalidHost(t *testing.T) {
	mp := MailProfile{
		Name:     "Test",
		UserName: "user",
		From:     "Test <test@example.com>",
		Host:     "invalidhost",
		Password: "password",
	}
	err := mp.SendMail("to@example.com", "Subject", "Body", nil)
	if err == nil {
		t.Error("expected error for invalid host, got nil")
	}
}

func TestSendMail_EmptyRecipient(t *testing.T) {
	mp := MailProfile{
		Name:     "Test",
		UserName: "user",
		From:     "Test <test@example.com>",
		Host:     "smtp.example.com:25",
		Password: "password",
	}
	err := mp.SendMail("", "Subject", "Body", nil)
	if err == nil {
		t.Error("Expected error for empty recipient, got nil")
	}
}

func TestSendMail_MultipleAttachments(t *testing.T) {
	mp := MailProfile{
		Name:     "Test",
		UserName: "user",
		From:     "Test <test@example.com>",
		Host:     "smtp.example.com:25",
		Password: "password",
	}
	attachments := []MailAttachment{
		{FileName: "a.txt", Data: []byte("This file is named a.txt")},
		{FileName: "b.txt", Data: []byte("This file is named b.txt")},
	}
	err := mp.SendMail("to@example.com", "Subject", "Body", attachments)
	if err == nil {
		t.Log("Multiple attachments sent (or at least, no error from gomail)")
	} else {
		t.Logf("SendMail returned error as expected (no real SMTP): %v", err)
	}
}

//change to actual Credentials
//userName, From, target, host
func TestSendMail_ValidHost(t *testing.T) {
	mp := MailProfile{
		Name:     "Jordan Phisher",
		UserName: "yourEmail.com",
		From:     "Jordan Phisher <yourEmail.com>",
		Host:     "smtp.gmail.com:587",
		Password: "",
	}
	target := "targetEmail.com"
	err := mp.SendTestMail(target)
	if err != nil {
		t.Error("Expected To pass, Got error\n &v", err)
	}
}

//change to actual Credentials
//userName, From, target, host
func TestSendMail_MultipleAttachments_valid(t *testing.T) {
	mp := MailProfile{
		Name:     "Jordan Phisher",
		UserName: "yourEmail.com",
		From:     "Jordan Phisher <yourEmail.com>",
		Host:     "smtp.gmail.com:587",
		Password: "",
	}
	target := "targetEmail.com"
	attachments := []MailAttachment{
		{FileName: "a.txt", Data: []byte("This file is named a.txt")},
		{FileName: "b.txt", Data: []byte("This file is named b.txt")},
	}
	err := mp.SendMail(target, "Attachment Test", "2 text files are attached", attachments)
	if err != nil {
		t.Error("Expected To pass, Got error\n &v", err)
	}
}

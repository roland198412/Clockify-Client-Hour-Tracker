package client

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewSMTPClient(host string, port int, username, password string) *SMTPClient {
	return &SMTPClient{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (c *SMTPClient) Send(from string, to []string, subject, body, contentType string) error {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	tlsConfig := &tls.Config{
		ServerName: c.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("TLS dial error: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, c.Host)
	if err != nil {
		return fmt.Errorf("SMTP client creation error: %w", err)
	}
	defer client.Quit()

	auth := smtp.PlainAuth("", c.Username, c.Password, c.Host)
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP auth error: %w", err)
	}

	if err := client.Mail(from); err != nil {
		return fmt.Errorf("SMTP MAIL FROM error: %w", err)
	}
	for _, addr := range to {
		if err := client.Rcpt(addr); err != nil {
			return fmt.Errorf("SMTP RCPT TO error for %s: %w", addr, err)
		}
	}

	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("SMTP DATA error: %w", err)
	}
	defer wc.Close()

	// Construct full message with headers
	msg := fmt.Sprintf("From: %s\r\n", from)
	msg += fmt.Sprintf("To: %s\r\n", to[0]) // you can join multiple recipients with comma if needed
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += fmt.Sprintf("MIME-Version: 1.0\r\n")
	msg += fmt.Sprintf("Content-Type: %s; charset=\"UTF-8\"\r\n", contentType)
	msg += "\r\n" // blank line between headers and body
	msg += body
	_, err = wc.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("write message error: %w", err)
	}

	return nil
}

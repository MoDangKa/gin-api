package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(to, resetURL string) error {
	gmailName := os.Getenv("GMAIL_Name")
	gmailUsername := os.Getenv("GMAIL_USERNAME")
	gmailPassword := os.Getenv("GMAIL_PASSWORD")

	m := gomail.NewMessage()

	m.SetHeader("From", fmt.Sprintf("%s <%s>", gmailName, gmailUsername))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Password Reset Request")

	emailBody := fmt.Sprintf("Hello,\n\nPlease click the following link to reset your password: %s\n\nBest regards,\n%s", resetURL, gmailName)
	m.SetBody("text/plain", emailBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, gmailUsername, gmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

package email

import (
	"log"
	"os"
	"strconv"

	"github.com/flambra/sender/internal/domain"
)

var (
	SMTP domain.SMTPConfig
)

func LoadSMTPConfig() {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM")

	if host == "" || portStr == "" || username == "" || password == "" || from == "" {
		log.Fatal("missing required SMTP environment variables")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("invalid SMTP port: %v", err)
	}

	SMTP = domain.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func GetSMTPConfig() domain.SMTPConfig  {
	return SMTP
}

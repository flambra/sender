package queue

import (
	"log"

	"github.com/flambra/sender/internal/config"
	"github.com/flambra/sender/internal/domain"
	"gopkg.in/gomail.v2"
)

var emailQueue = make(chan domain.EmailTask, 100)

func init() {
	go processEmailTask()
	log.Println("Email Queue initialized")
}

func EnqueueEmailTask(request domain.EmailRequest, template domain.EmailTemplate, body string) {
	emailQueue <- domain.EmailTask{
		Request:  request,
		Template: template,
		Body:     body,
	}
}

func processEmailTask() {
	for task := range emailQueue {
		sendEmail(task)
	}
}

func sendEmail(task domain.EmailTask) {
	SMTP := config.GetSMTPConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", SMTP.From)
	m.SetHeader("To", task.Request.To)
	m.SetHeader("Subject", task.Template.Subject)
	m.SetBody("text/html", task.Body)

	dialer := gomail.NewDialer(SMTP.Host, SMTP.Port, SMTP.Username, SMTP.Password)

	if err := dialer.DialAndSend(m); err != nil {
		log.Printf("failed to send Email to: %s with error: %v", task.Request.To, err)
	}
}

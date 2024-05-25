package template

import (
	"bytes"
	"html/template"

	"github.com/flambra/sender/internal/domain"
)

// Process replaces placeholders in the email template with actual values.
// It takes a template and recipient's name, then returns the processed HTML content.
//
//	func SendEmail() {
//		SMSTemplate := domain.SMSTemplate{
//			Content: "Hello, {{.RecipientName}}!",
//		}
//
//		processedContent, err := Process(SMSTemplate, "John Doe")
//		if err != nil {
//			log.Fatalf("Failed to process template: %v", err)
//		}
//		fmt.Println(processedContent)
//	}
//
// Hello, John Doe!
func Process(t domain.SMSTemplate, recipientName string) (string, error) {
	if recipientName == "" {
		return t.Text, nil
	}

	tmpl, err := template.New("emailTemplate").Parse(t.Text)
	if err != nil {
		return "", err
	}

	var text bytes.Buffer
	data := struct {
		RecipientName string
	}{
		RecipientName: recipientName,
	}

	if err := tmpl.Execute(&text, data); err != nil {
		return "", err
	}

	return text.String(), nil
}

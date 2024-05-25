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
//		templateEmail := domain.TemplateEmail{
//			Body: "<html><body>Welcome, {{.RecipientName}}!</body></html>",
//		}
// 		
//		processedContent, err := Process(templateEmail, "John Doe")
//		if err != nil {
//			log.Fatalf("Failed to process template: %v", err)
//		}
//		fmt.Println(processedContent)
//	}
//
// Welcome, John Doe!
func Process(t domain.TemplateEmail, recipientName string) (string, error) {
	tmpl, err := template.New("emailTemplate").Parse(t.Body)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	data := struct {
		RecipientName string
	}{
		RecipientName: recipientName,
	}

	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}

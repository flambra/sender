package template

import (
	"bytes"
	"html/template"

	"github.com/flambra/sender/internal/domain"
)

// ProcessTemplate processes the SMS template with the provided variables.
// It takes a template and a map of variables, then returns the processed text content.
func Process(t domain.SMSTemplate, variables map[string]interface{}) (string, error) {
	if len(variables) == 0 {
		return t.Message, nil
	}

	tmpl, err := template.New("smsTemplate").Parse(t.Message)
	if err != nil {
		return "", err
	}

	var content bytes.Buffer
	if err := tmpl.Execute(&content, variables); err != nil {
		return "", err
	}

	return content.String(), nil
}

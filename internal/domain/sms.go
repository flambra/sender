package domain


// SMSRequest represents the request structure for sending an SMS
type SMSRequest struct {
	To           string                 `json:"to"`
	TemplateName string                 `json:"template_name"`
	Variables    map[string]interface{} `json:"variables"` // Generalized to hold various template variables
}
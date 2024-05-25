package domain

type EmailRequest struct {
	To            string `json:"to"`
	TemplateName  string `json:"template_name"`
	Variables    map[string]interface{} `json:"variables"` // Generalized to hold various template variables
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

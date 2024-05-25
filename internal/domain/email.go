package domain

type EmailRequest struct {
	To           string `json:"to"`
	Subject      string `json:"subject"`
	TemplateName string `json:"template_name"`
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

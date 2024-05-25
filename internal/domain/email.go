package domain

type EmailRequest struct {
	To            string `json:"to"`
	TemplateName  string `json:"template_name"`
	RecipientName string `json:"recipient_name"`
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

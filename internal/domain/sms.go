package domain

type SmsRequest struct {
	To            string `json:"to"`
	TemplateName  string `json:"template_name"`
	RecipientName string `json:"recipient_name"`
}

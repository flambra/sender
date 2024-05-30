package domain

type EmailTask struct {
	Request  EmailRequest
	Template EmailTemplate
	Body     string
}

type SMSTask struct {
	Request  SMSRequest
	Template SMSTemplate
	Message  string
}

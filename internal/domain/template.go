package domain

import (
	"time"

	"gorm.io/gorm"
)

// EMAIL
type EmailTemplate struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Subject   string
	Body      string
}

type EmailTemplateCreateRequest struct {
	Name    string `form:"name"`
	Subject string `form:"subject"`
	Body    string `form:"body"`
}

type EmailTemplateUpdateRequest struct {
	Name    string `form:"name"`
	Subject string `form:"subject"`
	Body    string `form:"body"`
}

// SMS
type SMSTemplate struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Text      string
}

type SMSTemplateCreateRequest struct {
	Name string `form:"name"`
	Text string `form:"text"`
}

type SMSTemplateUpdateRequest struct {
	Name string `form:"name"`
	Text string `form:"text"`
}

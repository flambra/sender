package domain

import (
	"time"

	"gorm.io/gorm"
)

// EMAIL
type TemplateEmail struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Subject   string
	Body      string
}

type TemplateEmailCreateRequest struct {
	Name    string `form:"name"`
	Subject string `form:"subject"`
	Body    string `form:"body"`
}

type TemplateEmailUpdateRequest struct {
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

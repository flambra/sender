package domain

import (
	"time"

	"gorm.io/gorm"
)

type TemplateEmail struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Title     string
	Body      string
}

type TemplateEmailCreateRequest struct {
	Name  string `form:"name"`
	Title string `form:"title"`
	Body  string `form:"body"`
}

type TemplateEmailUpdateRequest struct {
	Name  string `form:"name"`
	Title string `form:"title"`
	Body  string `form:"body"`
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          int       `gorm:"primary_key" json:"id" validate:"gte=0"`
	Title       string    `gorm:"not null" json:"title" validate:"required"`
	Description string    `json:"description"`
	ExpiryDate  time.Time `gorm:"not null" json:"expiry_date" validate:"required"`
	Complete    int8      `gorm:"not null" json:"complete" validate:"gte=0,lte=100"`
}

func (t *Task) isComplete() bool {
	if t.Complete == 100 {
		return true
	} else {
		return false
	}
}

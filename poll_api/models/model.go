package models

import (
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	StartDate time.Time      `json:"start_date"`
	EndDate   time.Time      `json:"end_date"`
	Options   []PollOption   `json:"options"`
	Status    string         `json:"status"` 
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PollOption struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	PollID      uint           `json:"poll_id"`
	Description string         `json:"description"`
	Votes       int            `json:"votes"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
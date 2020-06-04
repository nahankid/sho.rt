package models

import (
	"time"
)

// Visit model
type Visit struct {
	APIKey    string    `gorm:"not null" json:"-"`
	LongURL   string    `gorm:"not null" json:"long_url"`
	ShortURL  string    `gorm:"unique;not null" json:"short_url"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	VisitedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"visited_at"`
}

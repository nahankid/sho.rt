package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// URL model
type URL struct {
	gorm.Model `json:"-"`
	APIKey     string    `gorm:"not null" json:"-"`
	LongURL    string    `gorm:"not null" json:"long_url"`
	ShortURL   string    `gorm:"unique_index;not null" json:"short_url"`
	ExpiresAt  time.Time `gorm:"not null;" json:"-"`
}

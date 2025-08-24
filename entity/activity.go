package entity

import (
	"time"
)

type Activity struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UserID    *string `gorm:"type:uuid;index"`
	Action    string  `gorm:"size:255"`
	Path      string  `gorm:"size:255"`
	Method    string  `gorm:"size:10"`
	IP        string  `gorm:"size:50"`
	UserAgent string  `gorm:"size:500"`
	Request   string  `gorm:"type:text"`
	Status    int     `gorm:"type:int"`
	CreatedAt time.Time
}

package model

import "time"

type Article struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Url       string    `json:"url" gorm:"not null"`
	Status    uint      `json:"status" gorm:"not null"`   // 0: 未読, 1: 読中, 2: 読了, 3: 保留
	Category  uint      `json:"category" gorm:"not null"` // 0: tech, 1: idea, 2: document, 3: release note, 99: else
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    string    `json:"user_id" gorm:"not null"`
}

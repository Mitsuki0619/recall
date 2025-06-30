package model

type Quiz struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Question string `json:"question" gorm:"not null"`
	Answer   string `json:"answer" gorm:"not null"`
	Type     uint   `json:"type" gorm:"not null"` // 0: 選択, 1: 正誤, 2: 記述
}

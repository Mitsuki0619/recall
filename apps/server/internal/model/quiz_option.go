package model

type QuizOption struct {
	QuizId string `json:"quiz_id" gorm:"primaryKey"`
	Quiz   Quiz   `json:"quiz" gorm:"foreignKey:QuizId;constraint:OnDelete:CASCADE"`
	Number uint   `json:"number" gorm:"primaryKey"`
	Value  string `json:"value" gorm:"not null"`
}

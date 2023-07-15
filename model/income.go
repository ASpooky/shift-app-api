package model

import "time"

type Incom struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	State     uint      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type IncomResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	State     uint      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

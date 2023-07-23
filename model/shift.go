package model

import "time"

type Shift struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartTime time.Time `json:"starttime" gorm:"not null"`
	EndTime   time.Time `json:"endtime" gorm:"not null"`
	UpdateAt  time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type ShiftResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

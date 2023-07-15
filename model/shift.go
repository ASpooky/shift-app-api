package model

import "time"

type Shift struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Date        time.Time `json:"attendance_day" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
	Workspace   Workspace `json:"workspace"`
	WorkspaceID uint      `json:"workspace_id" gorm:"not null"`
}

type ShiftResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Date      time.Time `json:"attendance_day" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

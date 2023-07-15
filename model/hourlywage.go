package model

import "time"

type HourlyWage struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
	H1          uint      `json:"h1"`
	H2          uint      `json:"h2"`
	H3          uint      `json:"h3"`
	H4          uint      `json:"h4"`
	H5          uint      `json:"h5"`
	H6          uint      `json:"h6"`
	H7          uint      `json:"h7"`
	H8          uint      `json:"h8"`
	H9          uint      `json:"h9"`
	H10         uint      `json:"h10"`
	H11         uint      `json:"h11"`
	H12         uint      `json:"h12"`
	H13         uint      `json:"h13"`
	H14         uint      `json:"h14"`
	H15         uint      `json:"h15"`
	H16         uint      `json:"h16"`
	H17         uint      `json:"h17"`
	H18         uint      `json:"h18"`
	H19         uint      `json:"h19"`
	H20         uint      `json:"h20"`
	H21         uint      `json:"h21"`
	H22         uint      `json:"h22"`
	H23         uint      `json:"h23"`
	H24         uint      `json:"h24"`
	Workspace   Workspace `json:"workspace" gorm:"foreignKey:WorkspaceId;constraint:OnDelete:CASCADE"`
	WorkspaceId uint      `json:"workspace_id" gorm:"not null"`
}

type HourlyWageResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	H1        uint      `json:"00:00~01:00"`
	H2        uint      `json:"01:00~02:00"`
	H3        uint      `json:"02:00~03:00"`
	H4        uint      `json:"03:00~04:00"`
	H5        uint      `json:"04:00~05:00"`
	H6        uint      `json:"05:00~06:00"`
	H7        uint      `json:"06:00~07:00"`
	H8        uint      `json:"07:00~08:00"`
	H9        uint      `json:"08:00~09:00"`
	H10       uint      `json:"09:00~10:00"`
	H11       uint      `json:"10:00~11:00"`
	H12       uint      `json:"11:00~12:00"`
	H13       uint      `json:"12:00~13:00"`
	H14       uint      `json:"13:00~14:00"`
	H15       uint      `json:"14:00~15:00"`
	H16       uint      `json:"15:00~16:00"`
	H17       uint      `json:"16:00~17:00"`
	H18       uint      `json:"17:00~18:00"`
	H19       uint      `json:"18:00~19:00"`
	H21       uint      `json:"19:00~20:00"`
	H22       uint      `json:"20:00~21:00"`
	H23       uint      `json:"21:00~22:00"`
	H24       uint      `json:"22:00~23:00"`
	H25       uint      `json:"23:00~24:00"`
}

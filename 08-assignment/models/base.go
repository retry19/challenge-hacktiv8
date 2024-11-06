package models

import "time"

type BaseModel struct {
	Id        uint      `json:"id,omitempty" gorm:"column:id;primary_key;auto_increment"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:createdAt;index"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updatedAt;index"`
}

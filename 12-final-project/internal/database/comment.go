package database

type Comment struct {
	BaseModel
	UserId  uint64 `json:"user_id" gorm:"column:user_id;not null"`
	PhotoId uint64 `json:"photo_id" gorm:"column:photo_id;not null" validate:"required,min=1"`
	Message string `json:"message" gorm:"column:message;not null" validate:"required,max=255"`
}

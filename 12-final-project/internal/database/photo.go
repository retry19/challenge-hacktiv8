package database

type Photo struct {
	BaseModel
	Title    string `json:"title" gorm:"column:title;not null" validate:"required,max=255"`
	Caption  string `json:"caption" gorm:"column:caption" validate:"max=255"`
	PhotoUrl string `json:"photo_url" gorm:"column:photo_url;not null" validate:"required,url,max=255"`
	UserId   uint64 `json:"user_id" gorm:"column:user_id;not null"`
}

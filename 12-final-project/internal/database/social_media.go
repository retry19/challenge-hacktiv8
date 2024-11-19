package database

type SocialMedia struct {
	BaseModel
	UserId         uint64 `json:"user_id" gorm:"column:user_id;not null"`
	Name           string `json:"name" gorm:"column:name;not null" validate:"required,max=255"`
	SocialMediaUrl string `json:"social_media_url" gorm:"column:social_media_url;not null" validate:"required,url,max=255"`
}

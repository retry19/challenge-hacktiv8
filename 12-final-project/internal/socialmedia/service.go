package socialmedia

import (
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"gorm.io/gorm"
)

type SocialMediaService struct {
	repository *SocialMediaRepository
}

func NewSocialMediaService(db *gorm.DB) *SocialMediaService {
	repository := NewSocialMediaRepository(db)

	return &SocialMediaService{
		repository: repository,
	}
}

func (sms *SocialMediaService) Create(socialMedia *database.SocialMedia) error {
	return sms.repository.db.Create(socialMedia).Error
}

func (sms *SocialMediaService) GetAll() ([]database.SocialMedia, error) {
	var socialMedias []database.SocialMedia
	err := sms.repository.db.Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (sms *SocialMediaService) GetOne(id uint64) (*database.SocialMedia, error) {
	var socialMedia database.SocialMedia
	err := sms.repository.db.First(&socialMedia, id).Error
	if err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func (sms *SocialMediaService) Delete(id uint64) error {
	return sms.repository.db.Delete(&database.SocialMedia{}, id).Error
}

func (sms *SocialMediaService) Update(socialMedia *database.SocialMedia) error {
	return sms.repository.db.Save(socialMedia).Error
}

package photo

import (
	"errors"

	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"gorm.io/gorm"
)

var (
	ErrPhotoNotFound = errors.New("photo not found")
)

type PhotoService struct {
	repository *PhotoRepository
}

func NewPhotoService(db *gorm.DB) *PhotoService {
	repository := NewPhotoRepository(db)

	return &PhotoService{
		repository: repository,
	}
}

func (ps *PhotoService) Create(photo *database.Photo) error {
	return ps.repository.db.Create(photo).Error
}

func (ps *PhotoService) GetAll() ([]database.Photo, error) {
	var photos []database.Photo
	err := ps.repository.db.Find(&photos).Error
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (ps *PhotoService) GetOne(id uint64) (*database.Photo, error) {
	var photo database.Photo
	err := ps.repository.db.First(&photo, id).Error
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

func (ps *PhotoService) Delete(id uint64) error {
	return ps.repository.db.Delete(&database.Photo{}, id).Error
}

func (ps *PhotoService) Update(photo *database.Photo) error {
	return ps.repository.db.Save(photo).Error
}

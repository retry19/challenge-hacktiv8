package comment

import (
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"gorm.io/gorm"
)

type CommentService struct {
	repository *CommentRepository
}

func NewCommentService(db *gorm.DB) *CommentService {
	repository := NewCommentRepository(db)

	return &CommentService{
		repository: repository,
	}
}

func (cs *CommentService) Create(comment *database.Comment) error {
	return cs.repository.db.Create(comment).Error
}

func (cs *CommentService) GetAll() ([]database.Comment, error) {
	var comments []database.Comment
	err := cs.repository.db.Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cs *CommentService) GetOne(id uint64) (*database.Comment, error) {
	var comment database.Comment
	err := cs.repository.db.First(&comment, id).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (cs *CommentService) Delete(id uint64) error {
	return cs.repository.db.Delete(&database.Comment{}, id).Error
}

func (cs *CommentService) Update(comment *database.Comment) error {
	return cs.repository.db.Save(comment).Error
}

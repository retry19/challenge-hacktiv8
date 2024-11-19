package auth

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/config"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/hasher"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUserIdNotFound        = errors.New("user id not found")
	ErrUserIdNotString       = errors.New("user id is not string")
	ErrUserIdFailedParse     = errors.New("failed to parse user id")
)

type AuthService struct {
	repository *AuthRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	repository := NewAuthRepository(db)

	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Register(user *database.User) error {
	var err error

	_, err = s.repository.FindByUsername(user.Username)
	if err == nil {
		return ErrUsernameAlreadyExists
	}

	_, err = s.repository.FindByEmail(user.Email)
	if err == nil {
		return ErrEmailAlreadyExists
	}

	hashedPassword, err := hasher.HashPassword(string(user.Password))
	if err != nil {
		return err
	}

	user.Password = database.Password(hashedPassword)

	return s.repository.Create(user)
}

func (s *AuthService) Login(payload *database.LoginPayload, user *database.User) error {
	var err error
	var existingUser *database.User

	if payload.Username != "" {
		existingUser, err = s.repository.FindByUsername(payload.Username)
	} else if payload.Email != "" {
		existingUser, err = s.repository.FindByEmail(payload.Email)
	}
	if err != nil {
		return ErrInvalidCredentials
	}

	err = hasher.ComparePassword(string(existingUser.Password), string(payload.Password))
	if err != nil {
		return ErrInvalidCredentials
	}

	existingUser.Password = ""

	*user = *existingUser

	return nil
}

func (s *AuthService) GenerateToken(user *database.User) (string, error) {
	return hasher.GenerateJwt(uint(user.Id), config.JwtSecret)
}

func (s *AuthService) GetUserId(c *fiber.Ctx) (uint64, error) {
	userIdStr := fmt.Sprintf("%v", c.Locals("UserId"))
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return 0, errors.Join(ErrUserIdFailedParse, err)
	}

	return userId, nil
}

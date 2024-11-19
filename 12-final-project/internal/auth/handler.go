package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/utility"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/validator"
)

type AuthHandler struct {
	service        *AuthService
	contextTimeout time.Duration
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{
		service:        service,
		contextTimeout: 10 * time.Second,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	bodyPayload := new(database.User)
	if err := c.BodyParser(bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	if errs := validator.ValidateSchema(bodyPayload); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utility.Response{
			Status:  false,
			Message: errs[0].Message,
			Errors:  errs,
		})
	}

	err := h.service.Register(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utility.Response{
		Status:  true,
		Data:    bodyPayload,
		Message: "User registered successfully",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	body := new(database.LoginPayload)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	if errs := validator.ValidateSchema(body); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utility.Response{
			Status:  false,
			Message: errs[0].Message,
			Errors:  errs,
		})
	}

	var user database.User
	err := h.service.Login(body, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	token, err := h.service.GenerateToken(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status: true,
		Data: struct {
			database.User
			Token string `json:"token"`
		}{
			User:  user,
			Token: token,
		},
		Message: "User logged in successfully",
	})
}

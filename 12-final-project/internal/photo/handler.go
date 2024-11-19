package photo

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/auth"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/utility"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/validator"
)

type PhotoHandler struct {
	service        *PhotoService
	authService    *auth.AuthService
	contextTimeout time.Duration
}

func NewPhotoHandler(photoService *PhotoService, authService *auth.AuthService) *PhotoHandler {
	return &PhotoHandler{
		service:        photoService,
		authService:    authService,
		contextTimeout: 10 * time.Second,
	}
}

func (ph *PhotoHandler) CreatePhoto(c *fiber.Ctx) error {
	bodyPayload := new(database.Photo)
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

	userId, err := ph.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	bodyPayload.UserId = userId

	err = ph.service.Create(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utility.Response{
		Status:  true,
		Data:    bodyPayload,
		Message: "Photo created successfully",
	})
}

func (ph *PhotoHandler) GetAll(c *fiber.Ctx) error {
	photos, err := ph.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Data:    photos,
		Message: "Photos retrieved successfully",
	})
}

func (ph *PhotoHandler) GetOne(c *fiber.Ctx) error {
	photoId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	photo, err := ph.service.GetOne(photoId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Data:    photo,
		Message: "Photo retrieved successfully",
	})
}

func (ph *PhotoHandler) DeletePhoto(c *fiber.Ctx) error {
	photoId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	userId, err := ph.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	photo, err := ph.service.GetOne(photoId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	if photo.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(utility.Response{
			Status:  false,
			Message: "You are not authorized to delete this photo",
		})
	}

	err = ph.service.Delete(photoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Message: "Photo deleted successfully",
	})
}

func (ph *PhotoHandler) UpdatePhoto(c *fiber.Ctx) error {
	photoId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	userId, err := ph.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	bodyPayload := new(database.Photo)
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

	photo, err := ph.service.GetOne(photoId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	if photo.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(utility.Response{
			Status:  false,
			Message: "You are not authorized to update this photo",
		})
	}

	bodyPayload.Id = photo.Id
	bodyPayload.UserId = userId

	err = ph.service.Update(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Data:    bodyPayload,
		Message: "Photo updated successfully",
	})
}

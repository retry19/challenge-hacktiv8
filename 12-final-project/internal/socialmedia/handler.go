package socialmedia

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/auth"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/utility"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/validator"
)

type SocialMediaHandler struct {
	service        *SocialMediaService
	authService    *auth.AuthService
	contextTimeout time.Duration
}

func NewSocialMediaHandler(socialMediaService *SocialMediaService, authService *auth.AuthService) *SocialMediaHandler {
	return &SocialMediaHandler{
		service:        socialMediaService,
		authService:    authService,
		contextTimeout: 10 * time.Second,
	}
}

func (smh *SocialMediaHandler) CreateSocialMedia(c *fiber.Ctx) error {
	bodyPayload := new(database.SocialMedia)
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

	userId, err := smh.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	bodyPayload.UserId = userId

	err = smh.service.Create(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utility.Response{
		Status:  true,
		Data:    bodyPayload,
		Message: "Social media created successfully",
	})
}

func (smh *SocialMediaHandler) GetAll(c *fiber.Ctx) error {
	socialMedias, err := smh.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Data:    socialMedias,
		Message: "Social medias retrieved successfully",
	})
}

func (smh *SocialMediaHandler) GetOne(c *fiber.Ctx) error {
	socialMediaId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	socialMedia, err := smh.service.GetOne(socialMediaId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Data:    socialMedia,
		Message: "Social media retrieved successfully",
	})
}

func (smh *SocialMediaHandler) DeleteSocialMedia(c *fiber.Ctx) error {
	socialMediaId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	userId, err := smh.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	socialMedia, err := smh.service.GetOne(socialMediaId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	if socialMedia.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(utility.Response{
			Status:  false,
			Message: "You are not authorized to delete this social media",
		})
	}

	err = smh.service.Delete(socialMediaId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Message: "Social media deleted successfully",
	})
}

func (smh *SocialMediaHandler) UpdateSocialMedia(c *fiber.Ctx) error {
	socialMediaId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	userId, err := smh.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	bodyPayload := new(database.SocialMedia)
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

	socialMedia, err := smh.service.GetOne(socialMediaId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	if socialMedia.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(utility.Response{
			Status:  false,
			Message: "You are not authorized to update this social media",
		})
	}

	bodyPayload.Id = socialMedia.Id
	bodyPayload.UserId = userId

	err = smh.service.Update(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utility.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utility.Response{
		Status:  true,
		Data:    bodyPayload,
		Message: "Social media updated successfully",
	})
}

package comment

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/auth"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/database"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/photo"
	"github.com/retry19/challenge-hacktiv8/12-final-project/pkg/validator"
)

type CommentHandler struct {
	service        *CommentService
	authService    *auth.AuthService
	photoService   *photo.PhotoService
	contextTimeout time.Duration
}

func NewCommentHandler(commentService *CommentService, authService *auth.AuthService, photoService *photo.PhotoService) *CommentHandler {
	return &CommentHandler{
		service:        commentService,
		authService:    authService,
		photoService:   photoService,
		contextTimeout: 10 * time.Second,
	}
}

func (ch *CommentHandler) CreateComment(c *fiber.Ctx) error {
	bodyPayload := new(database.Comment)
	if err := c.BodyParser(bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if errs := validator.ValidateSchema(bodyPayload); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": errs[0].Message,
			"errors":  errs,
		})
	}

	userId, err := ch.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	bodyPayload.UserId = userId

	err = ch.service.Create(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"data":    bodyPayload,
		"message": "Comment created successfully",
	})
}

func (ch *CommentHandler) GetAll(c *fiber.Ctx) error {
	comments, err := ch.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"data":    comments,
		"message": "Comments retrieved successfully",
	})
}

func (ch *CommentHandler) GetOne(c *fiber.Ctx) error {
	commentId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	comment, err := ch.service.GetOne(uint64(commentId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"data":    comment,
		"message": "Comment retrieved successfully",
	})
}

func (ch *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	commentId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	userId, err := ch.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	comment, err := ch.service.GetOne(commentId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	_, err = ch.photoService.GetOne(comment.PhotoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": photo.ErrPhotoNotFound,
		})
	}

	if comment.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to delete this comment",
		})
	}

	err = ch.service.Delete(commentId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Comment deleted successfully",
	})
}

func (ch *CommentHandler) UpdateComment(c *fiber.Ctx) error {
	commentId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	userId, err := ch.authService.GetUserId(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	bodyPayload := new(database.Comment)
	if err := c.BodyParser(bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if errs := validator.ValidateSchema(bodyPayload); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  false,
			"message": errs[0].Message,
			"errors":  errs,
		})
	}

	comment, err := ch.service.GetOne(commentId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if comment.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to update this comment",
		})
	}

	_, err = ch.photoService.GetOne(comment.PhotoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": photo.ErrPhotoNotFound,
		})
	}

	bodyPayload.Id = comment.Id
	bodyPayload.UserId = userId

	err = ch.service.Update(bodyPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"data":    bodyPayload,
		"message": "Comment updated successfully",
	})
}

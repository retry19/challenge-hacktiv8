package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/retry19/challenge-hacktiv8/09-gemini-ai/helpers"
	"github.com/retry19/challenge-hacktiv8/09-gemini-ai/models"
)

func ViewAskPage(c fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Ask me anything",
	}, "_layouts/main")
}

func AskGemini(question string, lastConversation *[]models.Conversation) (string, error) {
	url := fmt.Sprintf("%s/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s", helpers.GeminiBaseUrl, helpers.GeminiApiKey)
	payload := struct {
		Contents []models.GeminiContent `json:"contents"`
	}{
		Contents: []models.GeminiContent{},
	}

	if len(*lastConversation) > 0 {
		for _, conversation := range *lastConversation {
			payload.Contents = append(payload.Contents, models.GeminiContent{
				Role:  "user",
				Parts: []models.GeminiPart{{Text: conversation.Question}},
			})
			payload.Contents = append(payload.Contents, models.GeminiContent{
				Role:  "model",
				Parts: []models.GeminiPart{{Text: conversation.Answer}},
			})
		}
	}

	payload.Contents = append(payload.Contents, models.GeminiContent{
		Role:  "user",
		Parts: []models.GeminiPart{{Text: question}},
	})

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseBody models.GeminiGenerateContentResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.GetAnswer(), nil
}

func AskHandler(c fiber.Ctx) error {
	bodyPayload := new(models.Ask)

	if err := c.Bind().Body(bodyPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(bodyPayload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	var conversations []models.Conversation

	lastConversationCookie := c.Cookies(models.ConversationCookieKey)
	if lastConversationCookie != "" {
		err := json.Unmarshal([]byte(lastConversationCookie), &conversations)
		if err != nil {
			c.ClearCookie(models.ConversationCookieKey)
			return c.Status(fiber.StatusBadRequest).JSON(models.Response{
				Status:  false,
				Message: err.Error(),
			})
		}
	}

	answer, err := AskGemini(bodyPayload.Question, &conversations)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	conversation := models.Conversation{
		Question: bodyPayload.Question,
		Answer:   answer,
		Time:     time.Now(),
	}

	conversations = append(conversations, conversation)

	conversationsJSON, err := json.Marshal(conversations)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = models.ConversationCookieKey
	cookie.Value = string(conversationsJSON)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	return c.JSON(models.Response{
		Status: true,
		Data:   conversation,
	})
}

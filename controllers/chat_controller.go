package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safatanc/venti-ai/models"
	"github.com/safatanc/venti-ai/services/ai"
)

type ChatController struct {
	GeminiService *ai.GeminiService
}

func NewChatController(geminiService *ai.GeminiService) *ChatController {
	return &ChatController{
		GeminiService: geminiService,
	}
}

func (c *ChatController) HandleChat(ctx *fiber.Ctx) error {
	var chatRequest *models.ChatRequest
	if err := ctx.BodyParser(&chatRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	resp, err := c.GeminiService.GenerateText(chatRequest.Message)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(models.ChatResponse{
		Message: resp,
	})
}

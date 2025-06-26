package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/safatanc/venti-ai/models"
	"github.com/safatanc/venti-ai/services"
)

type ChatController struct {
	DeepseekService *services.OpenAIService
	GeminiService   *services.OpenAIService
}

func NewChatController(deepseekService *services.OpenAIService, geminiService *services.OpenAIService) *ChatController {
	return &ChatController{
		DeepseekService: deepseekService,
		GeminiService:   geminiService,
	}
}

func (c *ChatController) HandleChat(ctx *fiber.Ctx) error {
	// Upgrade to WebSocket
	return websocket.New(func(conn *websocket.Conn) {
		for {
			// Read message from client
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Parse chat request
			var chatRequest models.ChatRequest
			if err := json.Unmarshal(msg, &chatRequest); err != nil {
				conn.WriteJSON(models.ErrorResponse{
					Error: err.Error(),
				})
				continue
			}

			// Choose the appropriate service based on the model
			var service *services.OpenAIService
			switch chatRequest.Model {
			case models.DEEPSEEK_CHAT_MODEL:
				service = c.DeepseekService
			case models.GEMINI_FLASH_MODEL:
				service = c.GeminiService
			default:
				service = c.GeminiService
			}

			// Generate response
			if err := service.GenerateTextWS(conn, chatRequest.Message, chatRequest.SessionID); err != nil {
				conn.WriteJSON(models.ErrorResponse{
					Error: err.Error(),
				})
			}
		}
	})(ctx)
}

package main

import (
	"embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/safatanc/venti-ai/controllers"
	"github.com/safatanc/venti-ai/models"
	"github.com/safatanc/venti-ai/services"
	"github.com/safatanc/venti-ai/utils"
)

//go:embed data/*
var dataFS embed.FS

func main() {
	cfg := utils.LoadConfig()

	// Create services with embedded data
	deepseekService := services.NewOpenAIService(dataFS, models.GetModel(models.DEEPSEEK_CHAT_MODEL))
	geminiService := services.NewOpenAIService(dataFS, models.GetModel(models.GEMINI_FLASH_MODEL))

	// Create chat controller
	chatController := controllers.NewChatController(deepseekService, geminiService)

	app := fiber.New()

	// WebSocket middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Chat endpoint
	app.Get("/ws/chat", chatController.HandleChat)

	log.Fatal(app.Listen(cfg.GetServerAddress()))
}

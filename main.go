package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safatanc/venti-ai/controllers"
	"github.com/safatanc/venti-ai/services/ai"
	"github.com/safatanc/venti-ai/utils"
)

func main() {
	cfg := utils.LoadConfig()

	geminiService := ai.NewGeminiService("gemini-2.0-flash")
	chatController := controllers.NewChatController(geminiService)

	app := fiber.New()
	app.Post("/chat", chatController.HandleChat)

	app.Listen(cfg.GetServerAddress())
}

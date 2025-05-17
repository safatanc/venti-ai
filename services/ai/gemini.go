package ai

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

type GeminiService struct {
	Client *genai.Client
	Model  string
	Chat   *genai.Chat
}

func NewGeminiService(model string) *GeminiService {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	chat, err := client.Chats.Create(ctx, model, nil, nil)
	if err != nil {
		log.Fatalf("Failed to create chat: %v", err)
	}

	return &GeminiService{
		Client: client,
		Model:  model,
		Chat:   chat,
	}
}

func (s *GeminiService) GenerateText(prompt string) (string, error) {
	resp, err := s.Chat.SendMessage(context.Background(), genai.Part{
		Text: prompt,
	})
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}

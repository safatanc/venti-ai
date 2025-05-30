package services

import (
	"context"
	"embed"
	"log"
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/safatanc/venti-ai/models"
)

type OpenAIService struct {
	Client       *openai.Client
	Model        *models.Model
	systemPrompt string
	sessions     map[string]*models.ChatSession
	sessionMutex sync.RWMutex
}

func NewOpenAIService(dataFS embed.FS, model *models.Model) *OpenAIService {
	systemPrompt, err := dataFS.ReadFile("data/system_prompt.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if model.Name == "" {
		log.Fatal("Model name is not set")
		return nil
	}

	if model.APIKey == "" {
		log.Fatal("API key is not set for model " + model.Name)
		return nil
	}

	client := openai.NewClient(
		option.WithBaseURL(model.BaseURL),
		option.WithAPIKey(model.APIKey),
	)
	return &OpenAIService{
		Client:       &client,
		Model:        model,
		systemPrompt: string(systemPrompt),
		sessions:     make(map[string]*models.ChatSession),
	}
}

func (s *OpenAIService) getOrCreateSession(sessionID string) string {
	if sessionID == "" {
		sessionID = uuid.New().String()
	}

	s.sessionMutex.Lock()
	defer s.sessionMutex.Unlock()

	if _, exists := s.sessions[sessionID]; !exists {
		s.sessions[sessionID] = &models.ChatSession{
			Messages: []models.ChatMessage{
				{
					Role:    "system",
					Content: s.systemPrompt,
				},
			},
		}
	}

	return sessionID
}

func (s *OpenAIService) GenerateTextWS(conn *websocket.Conn, message string, sessionID string) error {
	start := time.Now()

	// Get or create session
	sessionID = s.getOrCreateSession(sessionID)

	// Get session messages
	s.sessionMutex.RLock()
	session := s.sessions[sessionID]
	messages := make([]models.ChatMessage, len(session.Messages))
	copy(messages, session.Messages)
	s.sessionMutex.RUnlock()

	// Add user message
	messages = append(messages, models.ChatMessage{
		Role:    "user",
		Content: message,
	})

	// Convert messages to OpenAI format
	openaiMessages := make([]openai.ChatCompletionMessageParamUnion, len(messages))
	for i, msg := range messages {
		switch msg.Role {
		case "system":
			openaiMessages[i] = openai.ChatCompletionMessageParamUnion{
				OfSystem: &openai.ChatCompletionSystemMessageParam{
					Role: "system",
					Content: openai.ChatCompletionSystemMessageParamContentUnion{
						OfString: openai.String(msg.Content),
					},
				},
			}
		case "user":
			openaiMessages[i] = openai.ChatCompletionMessageParamUnion{
				OfUser: &openai.ChatCompletionUserMessageParam{
					Role: "user",
					Content: openai.ChatCompletionUserMessageParamContentUnion{
						OfString: openai.String(msg.Content),
					},
				},
			}
		case "assistant":
			openaiMessages[i] = openai.ChatCompletionMessageParamUnion{
				OfAssistant: &openai.ChatCompletionAssistantMessageParam{
					Role: "assistant",
					Content: openai.ChatCompletionAssistantMessageParamContentUnion{
						OfString: openai.String(msg.Content),
					},
				},
			}
		}
	}

	// Generate response
	stream := s.Client.Chat.Completions.NewStreaming(context.Background(), openai.ChatCompletionNewParams{
		Model:    s.Model.Name,
		Messages: openaiMessages,
	})
	defer stream.Close()

	var fullContent string

	for stream.Next() {
		resp := stream.Current()
		if len(resp.Choices) > 0 {
			chunk := resp.Choices[0].Delta.Content
			if chunk != "" {
				fullContent += chunk
				response := models.ChatResponse{
					Model:        s.Model.Name,
					Message:      chunk,
					SessionID:    sessionID,
					ResponseTime: time.Since(start).Milliseconds(),
				}
				if err := conn.WriteJSON(response); err != nil {
					return err
				}
			}
		}
	}

	// Add assistant response to session
	s.sessionMutex.Lock()
	s.sessions[sessionID].Messages = append(s.sessions[sessionID].Messages, models.ChatMessage{
		Role:    "assistant",
		Content: fullContent,
	})
	s.sessionMutex.Unlock()

	// Send done message
	return conn.WriteJSON(models.ChatResponse{
		Model:        s.Model.Name,
		Message:      "[DONE]",
		SessionID:    sessionID,
		ResponseTime: time.Since(start).Milliseconds(),
	})
}

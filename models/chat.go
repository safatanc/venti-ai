package models

// ChatMessage represents a single message in a chat
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatSession represents a chat session with its messages
type ChatSession struct {
	Messages []ChatMessage
}

// ChatRequest represents a request to chat with the AI
type ChatRequest struct {
	Model                  string  `json:"model"`
	Message                string  `json:"message"`
	SessionID              string  `json:"session_id,omitempty"`
	AdditionalSystemPrompt *string `json:"additional_system_prompt,omitempty"`
}

// ChatResponse represents a response from the AI
type ChatResponse struct {
	Model        string `json:"model"`
	Message      string `json:"message"`
	SessionID    string `json:"session_id"`
	ResponseTime int64  `json:"response_time"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

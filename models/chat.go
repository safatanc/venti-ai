package models

// ChatRequest represents a request to chat with the AI
type ChatRequest struct {
	Message   string `json:"message"`
	SessionID string `json:"session_id,omitempty"`
}

// ChatResponse represents a response from the AI
type ChatResponse struct {
	Message   string `json:"message"`
	SessionID string `json:"session_id"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

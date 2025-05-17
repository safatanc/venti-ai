# Venti AI - Chatbot Service

Venti AI is a Go-based service that provides a REST API for interacting with various AI models. Currently, it supports Google's Gemini Flash AI.

## Features

- Service-based architecture, easily extendable to support multiple AI providers
- REST API for chat interactions
- Session management for persistent conversations
- Configurable settings via environment variables

## Requirements

- Go 1.24 or later
- Google Gemini API key

## Setup

1. Set the required environment variables:

```bash
export GEMINI_API_KEY="your-gemini-api-key"
# Optional configs:
export SERVER_PORT=8080
export AI_SERVICE_TYPE=gemini
export DEFAULT_LANGUAGE=en
```

2. Build and run the application:

```bash
go build
./venti-ai
```

## API Usage

### Chat Endpoint

**URL**: `/api/chat`
**Method**: `POST`
**Content-Type**: `application/json`

**Request Body**:
```json
{
  "message": "Hello, how are you?",
  "session_id": "optional-session-id"
}
```

**Response**:
```json
{
  "message": "I'm doing well, thank you for asking! How can I help you today?",
  "session_id": "session_123456789"
}
```

**Example curl request**:
```bash
curl -X POST http://localhost:8080/api/chat \
  -H "Content-Type: application/json" \
  -d '{"message": "Tell me a joke"}'
```

## Adding Support for Other AI Models

To add a new AI service:

1. Implement the AI service interface in a new file under `services/ai/`
2. Add the new service type to `factory.go`
3. Update `GetService()` to handle the new service type

## Architecture

- `services/ai/`: AI service interfaces and implementations
- `controllers/`: HTTP request handlers
- `models/`: Data structures for requests and responses
- `utils/`: Utility functions and configuration

## License

[MIT License](LICENSE) 
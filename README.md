# Venti AI - Chatbot Service

Venti AI is a Go-based service that provides WebSocket API for interacting with various AI models. Currently, it supports DeepSeek Chat and Google's Gemini Flash AI.

## Features

- WebSocket-based real-time chat with streaming responses
- Service-based architecture, easily extendable to support multiple AI providers
- Session management for persistent conversations
- Response time tracking
- Multiple AI model support
- Configurable settings via environment variables

## Requirements

- Go 1.24 or later
- DeepSeek API key
- Google Gemini API key

## Setup

1. Set the required environment variables:

```bash
export GEMINI_API_KEY="your-gemini-api-key"
export DEEPSEEK_API_KEY="your-deepseek-api-key"
# Optional configs:
export SERVER_PORT=8080
```

2. Build and run the application:

```bash
go build
./venti-ai
```

## API Usage

### WebSocket Chat Endpoint

**URL**: `ws://localhost:8080/ws/chat`
**Protocol**: `WebSocket`

**Request Message Format**:
```json
{
  "model": "deepseek-chat",
  "message": "Hello, how are you?",
  "session_id": "optional-session-id"
}
```

Available models:
- `deepseek-chat`: DeepSeek Chat model
- `gemini-2.0-flash`: Google Gemini Flash model

**Response Message Format**:
```json
{
  "model": "deepseek-chat",
  "message": "chunk of response",
  "session_id": "session_123456789",
  "response_time": 123
}
```

Special messages:
- `[DONE]` in message field indicates completion of response
- Error responses:
```json
{
  "error": "error message"
}
```

**Example JavaScript usage**:
```javascript
const ws = new WebSocket('ws://localhost:8080/ws/chat');

ws.onmessage = (event) => {
  const response = JSON.parse(event.data);
  
  if (response.message === '[DONE]') {
    console.log('Chat completed');
    return;
  }

  console.log('Model:', response.model);
  console.log('Message:', response.message);
  console.log('Session ID:', response.session_id);
  console.log('Response Time:', response.response_time, 'ms');
};

// Send a message
ws.send(JSON.stringify({
  model: 'deepseek-chat',
  message: 'Hello!',
  session_id: 'optional-session-id'
}));
```

## Adding Support for Other AI Models

To add a new AI model:

1. Add the model configuration in `models/model.go`
2. Update the OpenAI service to handle the new model
3. Update the chat controller to route requests to the new model

## Architecture

- `services/`: AI service implementations
- `controllers/`: WebSocket request handlers
- `models/`: Data structures and model configurations
- `utils/`: Utility functions and configuration
- `data/`: System prompts and other embedded data

## License

[MIT License](LICENSE) 
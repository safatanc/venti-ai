package models

import "os"

type Model struct {
	Name    string
	BaseURL string
	APIKey  string
}

const (
	DEEPSEEK_CHAT_MODEL = "deepseek-chat"
	GEMINI_FLASH_MODEL  = "gemini-2.0-flash"
)

func GetModel(name string) *Model {
	switch name {
	case DEEPSEEK_CHAT_MODEL:
		return &Model{
			Name:    DEEPSEEK_CHAT_MODEL,
			BaseURL: "https://api.deepseek.com/v1",
			APIKey:  os.Getenv("DEEPSEEK_API_KEY"),
		}
	case GEMINI_FLASH_MODEL:
		return &Model{
			Name:    GEMINI_FLASH_MODEL,
			BaseURL: "https://generativelanguage.googleapis.com/v1beta",
			APIKey:  os.Getenv("GEMINI_API_KEY"),
		}
	default:
		return nil
	}
}

func GetModelNames() []string {
	return []string{
		DEEPSEEK_CHAT_MODEL,
		GEMINI_FLASH_MODEL,
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	ollama "github.com/ollama/ollama/api"
	"github.com/sashabaranov/go-openai"
	"gosamples/utility/internal/application"
	"gosamples/utility/internal/infra"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalf("Error reading clipboard: %v", err)
	}
	ctx := context.Background()
	aiType := os.Getenv("AI_TYPE")
	var repo application.Repository
	switch aiType {
	case "openai":
		apiKey := os.Getenv("OPENAI_API_KEY")
		client := openai.NewClient(apiKey)
		repo = infra.NewOpenAIRepository(client, "gpt-4o-mini")
	case "ollama":
		ollamaClient, err := ollama.ClientFromEnvironment()
		if err != nil {
			log.Fatalf("Error creating Ollama client: %v", err)
		}
		repo = infra.NewOllamaRepository(ollamaClient, "mistral-nemo")
	default:
		log.Fatalf("Unknown AI type: %s", aiType)
	}

	result, err := repo.FixText(ctx, text)
	if err != nil {
		log.Fatalf("Error fixing text: %v", err)
	}
	fmt.Printf("ai type: %s result: %s\n", aiType, result)

	err = clipboard.WriteAll(result)
	if err != nil {
		log.Fatalf("Error writing to clipboard: %v", err)
	}
	return
}

package infra

import (
	"context"
	"time"

	ollama "github.com/ollama/ollama/api"
)

type OllamaRepository struct {
	client *ollama.Client
	model  string
}

func NewOllamaRepository(client *ollama.Client, model string) *OllamaRepository {
	return &OllamaRepository{client: client, model: model}
}

func (r *OllamaRepository) FixText(ctx context.Context, input string) (string, error) {
	var resp string
	err := r.client.Generate(context.Background(), &ollama.GenerateRequest{
		Model:     r.model,
		Prompt:    `Please correct any grammar mistakes and make this text little bit polite:` + input,
		System:    "You are a helpful assistant that corrects grammar and makes text polite without changing the meaning. Return just the answer without any additional comments",
		Stream:    new(bool),
		KeepAlive: &ollama.Duration{Duration: time.Hour},
	}, func(response ollama.GenerateResponse) error {
		resp = response.Response
		return nil
	})
	if err != nil {
		return "", err
	}
	return resp, nil
}

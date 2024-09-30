package infra

import (
	"fmt"

	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAIRepository struct {
	client *openai.Client
	model  string
}

func NewOpenAIRepository(client *openai.Client, model string) *OpenAIRepository {
	return &OpenAIRepository{client: client, model: model}
}

func (r *OpenAIRepository) FixText(ctx context.Context, input string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: r.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are a helpful assistant that corrects grammar and makes text polite.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Return me just the answer without additional text. Please correct any grammar mistakes and make this text little bit polite: %s", input),
			},
		},
	}
	resp, err := r.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("error during API call: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}
	return resp.Choices[0].Message.Content, nil
}

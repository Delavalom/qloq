package openai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func Generate(content string) string {
	apiKey, isEmpty := os.LookupEnv("API_KEY")
	if !isEmpty {
		fmt.Println("Error loading the API_KEY try qloq config to set it up")
		os.Exit(1)
	}
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		return "ChatCompletion error: " + err.Error()
	}

	return resp.Choices[0].Message.Content
}

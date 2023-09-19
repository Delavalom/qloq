package openai

import (
	"context"
	"log"

	configFile "github.com/Delavalom/qloq/pkg/config"
	openai "github.com/sashabaranov/go-openai"
)

func Generate(content string) string {
	apiKey, err := configFile.Retrieve()
	if err != nil {
		log.Fatalln("Error loading the API_KEY try qloq config to set it up")
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

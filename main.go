package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

func main() {

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
  For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Generate(strings.Join(args, " ")))
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
}

func Generate(content string) string {
	envs, err := godotenv.Read(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
		os.Exit(1)
	}
	client := openai.NewClient(envs["API_KEY"])
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

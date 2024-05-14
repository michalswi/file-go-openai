package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	apiKeys := os.Getenv("API_KEYS")
	if apiKeys == "" {
		log.Fatal("API_KEYS is not set")
	}

	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("ReadFile error: %v\n", err)
		return
	}

	openaiClient := openai.NewClient(apiKeys)
	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// https://pkg.go.dev/github.com/sashabaranov/go-openai@v1.23.0#pkg-constants
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(
						"%s: %s",
						os.Args[2],
						content),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	comment := fmt.Sprintf("ChatGPT's review about `%s` file:\n %s", filePath, resp.Choices[0].Message.Content)
	fmt.Println(comment)
}

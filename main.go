package main

import (
	"context"
	"flag"
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

	var filePath string
	var desc string

	saveToFile := flag.Bool("out", false, "Save review output to file")
	flag.StringVar(&filePath, "file", "", "file path")
	flag.StringVar(&desc, "desc", "", "question to openai model")
	flag.Parse()

	getOpenAIResponse(apiKeys, filePath, desc, saveToFile)

}

func getOpenAIResponse(apiKeys string, filePath string, desc string, saveToFile *bool) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("ReadFile error: %v\n", err)
		return
	}

	fmt.Println("OpenAI review started..")
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
						desc,
						content),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	if *saveToFile {
		println("Saving to file: ", filePath+"_review")
		err = os.WriteFile(filePath+"_review", []byte(resp.Choices[0].Message.Content), 0644)
		if err != nil {
			fmt.Printf("WriteFile error: %v\n", err)
			return
		}
		println("Saved to file..")
	} else {
		comment := fmt.Sprintf("ChatGPT's review about `%s` file:\n %s", filePath, resp.Choices[0].Message.Content)
		fmt.Println(comment)
	}
}

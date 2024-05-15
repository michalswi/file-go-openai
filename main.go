package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/michalswi/color"
	openai "github.com/sashabaranov/go-openai"
)

func main() {

	var filePath string
	var message string
	var saveToFile bool

	flag.StringVar(&filePath, "f", "", "Path to the file to be reviewed")
	flag.StringVar(&filePath, "file", "", "Path to the file to be reviewed")
	flag.StringVar(&message, "q", "", "Message to OpenAI model")
	flag.StringVar(&message, "message", "", "Message to OpenAI model")
	flag.BoolVar(&saveToFile, "o", false, "Save file's review output to a file")
	flag.BoolVar(&saveToFile, "out", false, "Save file's review output to a file")

	flag.Usage = func() {
		h := []string{
			"Options:",
			"  -f, --file <path>/<file>  Path to the file to be reviewed [required]",
			"  -m, --message <string>    Message to OpenAI model [required]",
			"  -o, --out                 Save file's review output to a file [optional]",
			"\n",
		}
		fmt.Fprintf(os.Stderr, "%s", strings.Join(h, "\n"))
	}
	flag.Parse()

	apiKeys := os.Getenv("API_KEYS")
	if apiKeys == "" {
		log.Fatal("Please set the API_KEYS environment variable to your OpenAI API key")
	}

	resp, err := getOpenAIResponse(apiKeys, filePath, message)
	if err != nil {
		log.Fatalf("OpenAI review failed: %v\n", err)
	}

	if saveToFile {
		writeReview(resp, filePath)
	} else {
		fmt.Println(color.Format(color.GREEN, fmt.Sprintf("ChatGPT's review base on `%s` file:", filePath)))
		fmt.Println(color.Format(color.YELLOW, resp.Choices[0].Message.Content))
	}

}

// getOpenAIResponse reads the file at the given path, sends its content to the OpenAI API
// along with the provided question, and returns the API's response.
func getOpenAIResponse(apiKeys string, filePath string, message string) (resp openai.ChatCompletionResponse, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file: %v\n", err)
		return
	}

	fmt.Println(color.Format(color.GREEN, "OpenAI review started.."))
	openaiClient := openai.NewClient(apiKeys)

	resp, err = openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// https://pkg.go.dev/github.com/sashabaranov/go-openai@v1.23.0#pkg-constants
			Model: openai.GPT4,
			// Model: openai.GPT4Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(
						"%s: %s",
						message,
						content),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatGPT review failed: %v\n", err)
		return
	}

	return resp, nil

}

// writeReview checks if a file with the same name already exists. If it does, it asks the user
// if they want to overwrite it. If the user agrees or if the file doesn't exist, it writes the
// OpenAI response to the file.
func writeReview(resp openai.ChatCompletionResponse, filePath string) {
	reader := bufio.NewReader(os.Stdin)
	reviewFile := filePath + "_review"
	fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Saving review to file: %s", reviewFile)))
	_, err := os.Stat(reviewFile)
	if err == nil {
		fmt.Printf("File %s already exist, overwrite it? [Y,n]: ", reviewFile)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		switch command {
		case "n":
			fmt.Printf("Review not saved to file: %s\n", reviewFile)
			return
		case "N":
			fmt.Printf("Review not saved to file: %s\n", reviewFile)
			return
		case "y":
			writeToFile(reviewFile, resp.Choices[0].Message.Content)
			fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Review saved to file: %s", reviewFile)))
		case "Y":
			writeToFile(reviewFile, resp.Choices[0].Message.Content)
			fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Review saved to file: %s", reviewFile)))
		default:
			writeToFile(reviewFile, resp.Choices[0].Message.Content)
			fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Review saved to file: %s", reviewFile)))
		}
	} else {
		writeToFile(reviewFile, resp.Choices[0].Message.Content)
		fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Review saved to file: %s", reviewFile)))
	}
}

// writeToFile writes the provided content to a file at the given path. If an error occurs,
// it logs the error and returns.
func writeToFile(reviewFile string, content string) {
	err := os.WriteFile(reviewFile, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Write review to file failed: %v\n", err)
		return
	}
}

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/michalswi/color"
	openai "github.com/sashabaranov/go-openai"
)

const (
	patternURL = "https://raw.githubusercontent.com/michalswi/file-go-openai/main/patterns/"
	// patternURL = "https://raw.githubusercontent.com/michalswi/file-go-openai/dev/patterns/"
	patternFile = "pattern"
	// https://pkg.go.dev/github.com/sashabaranov/go-openai#pkg-constants
	openAImodel = openai.O1Mini
	// openAImodel = openai.GPT4oMini
	reviewFileExt = "_rev"
	filePerm      = 0644
)

func main() {

	var filePath string
	var message string
	var pattern string
	var saveToFile bool
	var inputQuery string
	var oaiVersion bool

	flag.StringVar(&filePath, "f", "", "Path to the file to be reviewed [required]")
	flag.StringVar(&filePath, "file", "", "Path to the file to be reviewed [required]")
	flag.StringVar(&message, "m", "", "Message to OpenAI model [required OR use '-p']")
	flag.StringVar(&message, "message", "", "Message to OpenAI model [required OR use '-p']")
	flag.StringVar(&pattern, "p", "", "Pattern name")
	flag.StringVar(&pattern, "pattern", "", "Pattern name")
	flag.BoolVar(&saveToFile, "o", false, "Save file's review output to a file [optional]")
	flag.BoolVar(&saveToFile, "out", false, "Save file's review output to a file [optional]")
	flag.BoolVar(&oaiVersion, "v", false, "Display OpenAI model version")
	flag.BoolVar(&oaiVersion, "version", false, "Display OpenAI model version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if oaiVersion {
		fmt.Println(openAImodel)
		os.Exit(0)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: The API_KEY env variable (to your OpenAI API account) is not set.")
		os.Exit(1)
	}

	if filePath == "" {
		fmt.Fprintln(os.Stderr, "Error: The file path must be specified using the -f or --file flag.")
		flag.Usage()
		os.Exit(1)
	}

	if message == "" && pattern == "" {
		fmt.Fprintln(os.Stderr, "Error: A message or pattern must be specified using the -m/--message or -p/--pattern flag.")
		flag.Usage()
		os.Exit(1)
	} else {
		if message != "" && pattern == "" {
			inputQuery = message
		}
		if message == "" && pattern != "" {
			resp, err := getPattern(pattern)
			if err != nil {
				log.Fatalf("Couldn't get pattern: %v\n", err)
			}
			inputQuery = resp
		}
	}

	resp, err := getOpenAIResponse(apiKey, filePath, inputQuery)
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

// getOpenAIResponse reads the content of the file at filePath and sends it
// along with the message to the OpenAI API. It returns the API's response.
func getOpenAIResponse(apiKey string, filePath string, message string) (resp openai.ChatCompletionResponse, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read file: %v\n", err)
		return
	}

	fmt.Println(color.Format(color.GREEN, "OpenAI review started.."))
	openaiClient := openai.NewClient(apiKey)

	resp, err = openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openAImodel,
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
		// fmt.Printf("ChatGPT review failed: %v\n", err)
		return
	}

	return resp, nil

}

// writeReview checks if a file with the same name already exists. If it does, it asks the user
// if they want to overwrite it. If the user agrees or if the file doesn't exist, it writes the
// OpenAI response to the file.
func writeReview(resp openai.ChatCompletionResponse, filePath string) {
	reader := bufio.NewReader(os.Stdin)
	reviewFile := filePath + reviewFileExt
	fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Saving review to file: %s", reviewFile)))
	_, err := os.Stat(reviewFile)

	// file exists
	if err == nil {
		fmt.Printf("File %s already exist, overwrite it? [Y,n]: ", reviewFile)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		if command == "n" || command == "no" {
			fmt.Printf("Review not saved to file: %s\n", reviewFile)
			return
		}
	}
	if err := writeToFile(reviewFile, resp.Choices[0].Message.Content); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		return
	}
	fmt.Println(color.Format(color.GREEN, fmt.Sprintf("Review saved to file: %s", reviewFile)))
}

// writeToFile writes the provided content to the specified file path.
func writeToFile(reviewFile string, content string) error {
	err := os.WriteFile(reviewFile, []byte(content), filePerm)
	if err != nil {
		log.Fatalf("Write review to file failed: %v\n", err)
	}
	return err
}

// getPattern is a function that retrieves a specific pattern from a remote server.
// It takes a single argument, patternName, which is a string representing the name of the pattern to retrieve.
func getPattern(patternName string) (string, error) {
	url := fmt.Sprintf("%s%s/%s", patternURL, patternName, patternFile)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch pattern from URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	return string(data), nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()
	llm, err := openai.New(openai.WithToken(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	completion, err := llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "What is the capital of France?"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}

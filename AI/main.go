package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	// Načíst .env soubor (z kořenového adresáře projektu)
	err := godotenv.Load("../.env")
	if err != nil {
		// Pokud .env neexistuje, pokračuj (možná je API key v system env)
		log.Println("Warning: .env file not found, using system environment variables")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}

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

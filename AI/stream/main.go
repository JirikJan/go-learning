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
	err := godotenv.Load("../../.env")
	if err != nil {
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

	// Příprava zpráv
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "Proč se vyplatí psát AI aplikace v Go?"),
	}

	// Streamování odpovědi pomocí callback funkce
	_, err = llm.GenerateContent(ctx, messages,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println() // Nový řádek na konci
}

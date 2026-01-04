package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/iamwavecut/go-tavily"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		// Pokud .env neexistuje, pokračuj (možná je API key v system env)
		log.Println("Warning: .env file not found, using system environment variables")
	}

	apiKey := os.Getenv("TAVILY_API_KEY")
	if apiKey == "" {
		log.Fatal("TAVILY_API_KEY environment variable is not set")
	}

	// Inicializace Tavily klienta a načtení API klíče z .env souboru
	client := tavily.New(apiKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// Zobrazení výstupů z vyhledávání
	result, err := client.SearchSimple(ctx, "Go programming language")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

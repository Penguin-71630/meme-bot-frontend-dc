package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Penguin-71630/meme-bot-frontend-dc/bot"
	"github.com/Penguin-71630/meme-bot-frontend-dc/config"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize bot
	discordBot, err := bot.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Start bot
	if err := discordBot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	// Wait for interrupt signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanup
	discordBot.Stop()
	fmt.Println("Bot stopped gracefully.")
}

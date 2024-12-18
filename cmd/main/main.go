package main

import (
	"log"
	"os"
	"telegram_invest_bot/internal/telegram"
)

func main() {
	bot, err := telegram.NewBot()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
		os.Exit(1)
	}

	if err := bot.RunBot(); err != nil {
		log.Fatalf("Failed to run bot: %v", err)
		os.Exit(1)
	}
}

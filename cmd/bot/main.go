package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/vhalmd/nomi-go-sdk"
	"log/slog"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	key := os.Getenv("NOMI_API_KEY")
	if key == "" {
		slog.Error("Nomi API key not found. Set NOMI_API_KEY on the .env file")
		os.Exit(1)
	}

	nomiId := os.Getenv("NOMI_ID")
	if key == "" {
		slog.Error("Nomi ID not found. Set NOMI_ID on the .env file")
		os.Exit(1)
	}

	telegramBotKey := os.Getenv("TELEGRAM_BOT_KEY")
	if key == "" {
		slog.Error("Telegram bot key not found. Set TELEGRAM_BOT_KEY on the .env file")
		os.Exit(1)
	}

	client := nomi.NewClient(key)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler(client, nomiId)),
	}

	b, err := bot.New(telegramBotKey, opts...)
	if err != nil {
		panic(err)
	}

	slog.Info("Starting bot...")
	b.Start(ctx)
}

func handler(client nomi.API, nomiId string) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		slog.Info("Received a new message. Repassing message to Nomi.", "channel", update.Message.Chat.Username, "content", update.Message.Text)
		reply, _ := client.SendMessage(
			nomiId,
			nomi.SendMessageBody{
				MessageText: "*from telegram* " + update.Message.Text,
			},
		)
		slog.Info("Nomi replied", "text", reply.ReplyMessage.Text)

		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   reply.ReplyMessage.Text,
		})
	}
}

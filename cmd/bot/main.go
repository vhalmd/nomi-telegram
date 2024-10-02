package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/vhalmd/nomi-go-sdk"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type BotConfig struct {
	NomiID        string
	TelegramToken string
	BotName       string
}

func main() {
	nomiKey := getEnvVarOrExit("NOMI_API_KEY", "Nomi API key")
	nomiIds := getEnvVarOrExit("NOMI_ID", "Nomi ID")
	nomiNames := getEnvVarOrExit("NOMI_NAME", "Nomi Name")
	telegramBotTokens := getEnvVarOrExit("TELEGRAM_BOT_TOKEN", "Telegram bot token")
	prefix := os.Getenv("PREFIX_MESSAGES_WITH")

	client := nomi.NewClient(nomiKey)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	parsedTokens := strings.Split(telegramBotTokens, ",")
	parsedNomiIds := strings.Split(nomiIds, ",")
	parsedNomiNames := strings.Split(nomiNames, ",")

	validateTokenAndIdCount(parsedTokens, parsedNomiIds, parsedNomiNames)
	botConfigs := parseConfig(parsedTokens, parsedNomiIds, parsedNomiNames)

	if len(parsedTokens) != len(parsedNomiIds) {
		fmt.Printf(
			"Nomi ID and Telegram Token counts don't match. Make sure you have one Nomi ID for each Telegram Token, and the order that they're configured should also match.\nFor example:\nNOMI_IDS=NOMI_1_ID,NOMI_1_ID\nTELEGRAM_BOT_TOKENS=BOT_TOKEN_1,BOT_TOKEN_2\n",
		)
		os.Exit(1)
	}

	for i, config := range botConfigs {
		slog.Info("Starting bot", "bot_name", config.BotName, "telegram_token", config.TelegramToken, "nomi_id", parsedNomiIds[i])
		go startBot(ctx, client, config, prefix)
	}

	<-ctx.Done()
}

func startBot(ctx context.Context, client nomi.API, cfg BotConfig, prefix string) {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler(client, cfg, prefix)),
	}

	b, err := bot.New(cfg.TelegramToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(client nomi.API, cfg BotConfig, prefix string) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message.Text == "/start" {
			slog.Info(fmt.Sprintf("[%s] received /start command. Ignoring message...", strings.ToUpper(cfg.BotName)))
			return
		}
		slog.Info(
			fmt.Sprintf("[%s] received a new message. Repassing message to Nomi.", strings.ToUpper(cfg.BotName)),
			"from", update.Message.Chat.Username,
			"content", update.Message.Text,
		)

		msg := ""
		if prefix != "" {
			msg += prefix + " "
		}
		msg += update.Message.Text

		reply, _ := client.SendMessage(
			cfg.NomiID,
			nomi.SendMessageBody{
				MessageText: msg,
			},
		)
		slog.Info(
			fmt.Sprintf("[%s] replied", strings.ToUpper(cfg.BotName)),
			"text", reply.ReplyMessage.Text,
		)

		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   reply.ReplyMessage.Text,
		})
	}
}

func getEnvVarOrExit(envVar, description string) string {
	value := os.Getenv(envVar)
	if value == "" {
		slog.Error(fmt.Sprintf("%s not found. Set %s in the .env file", description, envVar))
		os.Exit(1)
	}
	return value
}

func validateTokenAndIdCount(tokens, ids, names []string) {
	if len(tokens) != len(ids) || len(ids) != len(names) {
		fmt.Println("Mismatch between Nomi IDs, Telegram Tokens, and Nomi Names. Ensure one Nomi ID, one Telegram Token, and one Nomi Name per bot.")
		fmt.Println("Example configuration:")
		fmt.Println("NOMI_ID=NOMI_1_ID,NOMI_2_ID")
		fmt.Println("TELEGRAM_BOT_TOKEN=BOT_TOKEN_1,BOT_TOKEN_2")
		fmt.Println("NOMI_NAME=NOMI_NAME_1,NOMI_NAME_2")
		os.Exit(1)
	}
}

func parseConfig(tokens, nomiIds, nomiNames []string) []BotConfig {
	botConfigs := make([]BotConfig, len(nomiIds))
	for i := range nomiIds {
		botConfigs[i] = BotConfig{
			NomiID:        nomiIds[i],
			TelegramToken: tokens[i],
			BotName:       nomiNames[i],
		}
	}
	return botConfigs
}

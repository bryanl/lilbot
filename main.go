package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/bryanl/lilbot/internal/log"
	"github.com/bryanl/lilbot/pkg/bot"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config := Config{}

	flag.StringVar(&config.Token, "token", envString("LILBOT_TOKEN", ""), "bot token")
	flag.Parse()

	ctx := log.WithLogger(context.Background())

	logger := log.From(ctx)

	if err := run(ctx, config); err != nil {
		logger.Error(err, "run application")
		os.Exit(1)
	}
}

// Config is configuration for lilbot
type Config struct {
	// Token is the bot token
	Token string
}

func run(ctx context.Context, config Config) error {
	logger := log.From(ctx)

	brain := bot.NewBrain(ctx)

	logger.Info("creating session")
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return fmt.Errorf("create discord session: %w", err)
	}

	if err := brain.AddHandlers(session); err != nil {
		return fmt.Errorf("add handlers for session: %w", err)
	}

	if err := session.Open(); err != nil {
		return fmt.Errorf("open session: %w", err)
	}

	defer func() {
		if cErr := session.Close(); cErr != nil {
			logger.Error(err, "close session")
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	logger.Info("bot is ready to process handlers")
	<-stop
	logger.Info("stopping")

	return nil
}

func envString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

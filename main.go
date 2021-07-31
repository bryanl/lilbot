package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/bryanl/lilutil/goutil"
	"github.com/bryanl/lilutil/log"
	"github.com/bwmarrin/discordgo"

	"github.com/bryanl/lilbot/pkg/bot"
)

func main() {
	config := Config{}

	flag.StringVar(&config.Token, "token", envString("LILBOT_TOKEN", ""), "bot token")
	flag.StringVar(
		&config.PluginHostAddr,
		"plugin-host-addr",
		envString("LILBOT_PLUGIN_HOST_ADDR", ":50051"),
		"plugin host address")
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
	// PluginHostAddr is the address the plugin host server should listen on.
	PluginHostAddr string
}

func run(ctx context.Context, config Config) error {
	logger := log.From(ctx).WithName("main")

	runCtx, runCancel := context.WithCancel(ctx)
	defer runCancel()

	brain := bot.NewBrain(runCtx)

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

	if err := brain.CreateCommands(session); err != nil {
		return fmt.Errorf("create commands: %w", err)
	}

	pluginHost := bot.NewPluginHost()

	grpcServer, err := bot.NewGRPCServer(config.PluginHostAddr, pluginHost)
	if err != nil {
		return fmt.Errorf("create GRPC server: %w", err)
	}

	grpcServerDone, err := grpcServer.Start(runCtx)
	if err != nil {
		return fmt.Errorf("start GRPC server: %w", err)
	}

	goutil.HandleGracefulClose(ctx, runCancel, grpcServerDone)

	return nil
}

func envString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

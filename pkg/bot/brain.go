package bot

import (
	"context"
	"errors"
	"fmt"

	"github.com/bryanl/lilutil/log"
	"github.com/bwmarrin/discordgo"
	"github.com/go-logr/logr"
)

// Brain is a bot interface.
type Brain interface {
	// CreateCommand creates a command.
	CreateCommand(ctx context.Context, command *discordgo.ApplicationCommand) error
}

// DiscordBot is a live implementation of Brain.
type DiscordBot struct {
	session *discordgo.Session
	logger  logr.Logger
}

var _ Brain = &DiscordBot{}

// New creates an instance of DiscordBot.
func New(ctx context.Context, token string) (*DiscordBot, error) {
	logger := log.From(ctx).WithName("lilbot.DiscordBot")

	logger.Info("creating session")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("create discord session: %w", err)
	}

	bot := &DiscordBot{
		session: session,
		logger:  logger,
	}

	return bot, nil
}

// Start starts the bot.
func (bot *DiscordBot) Start(ctx context.Context) (<-chan struct{}, error) {
	bot.logger.Info("starting bot")

	if err := bot.addHandlers(bot.session); err != nil {
		return nil, fmt.Errorf("add handlers for session :%w", err)
	}

	if err := bot.session.Open(); err != nil {
		return nil, fmt.Errorf("open session: %w", err)
	}

	ch := make(chan struct{}, 1)

	go func() {
		<-ctx.Done()
		bot.logger.Info("closing discord session")

		if cErr := bot.session.Close(); cErr != nil {
			bot.logger.Error(cErr, "close discord session")
		}

		bot.logger.Info("session has closed")
		close(ch)
	}()

	return ch, nil
}

// CreateCommand creates a bot command.
func (bot *DiscordBot) CreateCommand(ctx context.Context, command *discordgo.ApplicationCommand) error {

	// TODO: store the output somewhere
	// _, err := session.ApplicationCommandCreate(session.State.User.ID, "", applicationCommand)
	// if err != nil {
	// 	return fmt.Errorf("create command '%s': %w", name, err)
	// }

	return nil
}

// Ready is a discord ready handler.
func (bot *DiscordBot) Ready(_ *discordgo.Session, ready *discordgo.Ready) {
	bot.logger.Info("bot is up", "ready", ready)
}

// ReceiveMessage is a discord receive message handler.
func (bot *DiscordBot) ReceiveMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if bot.isSelf(session, message) {
		return
	}

	if message.Content == "ping" {
		bot.speak(session, message.ChannelID, "Pong!")
	}

	if message.Content == "pong" {
		bot.speak(session, message.ChannelID, "Ping!")
	}

	if message.Content == "secret" {
		bot.dm(session, message.Author.ID, "A secret message")
	}

	bot.logger.Info("received message", "message", message)
}

// InteractionCreate creates an interaction.
func (bot *DiscordBot) InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

func (bot *DiscordBot) addHandlers(session *discordgo.Session) error {
	if session == nil {
		return errors.New("session is nil")
	}

	session.AddHandler(bot.Ready)
	session.AddHandler(bot.ReceiveMessage)
	session.AddHandler(bot.InteractionCreate)

	return nil
}

func (bot *DiscordBot) isSelf(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	if session == nil || message == nil {
		return false
	}

	if session.State.User.ID == message.Author.ID {
		return true
	}

	return false
}

func (bot *DiscordBot) dm(session *discordgo.Session, recipientID, message string) {
	channel, err := session.UserChannelCreate(recipientID)
	if err != nil {
		bot.logger.Error(err, "unable to create user channel", "recipientID", recipientID)
		return
	}

	bot.speak(session, channel.ID, message)
}

func (bot *DiscordBot) speak(session *discordgo.Session, channelID, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		bot.logger.Error(err, "send message")
	}
}

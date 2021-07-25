package bot

import (
	"context"
	"errors"

	"github.com/bryanl/lilbot/internal/log"
	"github.com/bwmarrin/discordgo"
	"github.com/go-logr/logr"
)

// Brain is the bot control center.
type Brain struct {
	logger logr.Logger
}

// NewBrain creates a new instance of Brain.
func NewBrain(ctx context.Context) *Brain {
	logger := log.From(ctx)

	bot := &Brain{
		logger: logger,
	}

	return bot
}

// AddHandlers adds handlers for a session.
func (brain *Brain) AddHandlers(session *discordgo.Session) error {
	if session == nil {
		return errors.New("session is nil")
	}

	session.AddHandler(brain.Ready)
	session.AddHandler(brain.ReceiveMessage)

	return nil
}

// Ready is a discord ready handler.
func (brain *Brain) Ready(_ *discordgo.Session, ready *discordgo.Ready) {
	brain.logger.Info("bot is up", "ready", ready)
}

// ReceiveMessage is a discord receive message handler.
func (brain *Brain) ReceiveMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if brain.isSelf(session, message) {
		return
	}

	if message.Content == "ping" {
		brain.speak(session, message.ChannelID, "Pong!")
	}

	if message.Content == "pong" {
		brain.speak(session, message.ChannelID, "Ping!")
	}

	if message.Content == "secret" {
		brain.dm(session, message.Author.ID, "A secret message")
	}

	brain.logger.Info("received message", "message", message)
}

func (brain *Brain) isSelf(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	if session == nil || message == nil {
		return false
	}

	if session.State.User.ID == message.Author.ID {
		return true
	}

	return false
}

func (brain *Brain) dm(session *discordgo.Session, recipientID, message string) {
	channel, err := session.UserChannelCreate(recipientID)
	if err != nil {
		brain.logger.Error(err, "unable to create user channel", "recipientID", recipientID)
		return
	}

	brain.speak(session, channel.ID, message)
}

func (brain *Brain) speak(session *discordgo.Session, channelID, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		brain.logger.Error(err, "send message")
	}
}

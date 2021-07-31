package bot

import (
	"context"
	"errors"
	"fmt"

	pluginostv1alpha1 "github.com/bryanl/lilbot/gen/go/pluginhost/v1alpha1"
	"github.com/bryanl/lilutil/log"
	"github.com/bwmarrin/discordgo"
)

// PluginHost is a plugin host interface.
type PluginHost interface {
	Register(ctx context.Context, request *pluginostv1alpha1.RegisterRequest) error
}

// LivePluginHost is a live implementation of PluginHost.
type LivePluginHost struct {
	brain Brain
}

var _ PluginHost = &LivePluginHost{}

// NewPluginHost creates an instance of LivePluginHost.
func NewPluginHost(bot Brain) *LivePluginHost {
	pluginHost := &LivePluginHost{
		brain: bot,
	}

	return pluginHost
}

// Register registers a plugin with the host.
func (pluginHost *LivePluginHost) Register(ctx context.Context, registerRequest *pluginostv1alpha1.RegisterRequest) error {
	logger := log.From(ctx)

	if registerRequest == nil {
		return errors.New("request is nil")
	}

	for actionName, action := range registerRequest.Actions {
		name := fmt.Sprintf("%s.%s", registerRequest.GetName(), actionName)
		logger.Info("registering command", "name", name)

		applicationCommand := &discordgo.ApplicationCommand{
			Name:        name,
			Description: action.GetDescription(),
			Options:     []*discordgo.ApplicationCommandOption{},
		}

		if err := pluginHost.brain.CreateCommand(ctx, applicationCommand); err != nil {
			return fmt.Errorf("create command '%s': %w", name, err)
		}
	}

	return nil
}

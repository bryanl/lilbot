package bot

import (
	"context"
	"fmt"

	"github.com/bryanl/lilutil/log"
)

// PluginHost is a plugin host interface.
type PluginHost interface {
	Register(ctx context.Context, request Command) error
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
func (pluginHost *LivePluginHost) Register(ctx context.Context, registerRequest Command) error {
	logger := log.From(ctx)

	for actionName, action := range registerRequest.Actions {
		name := fmt.Sprintf("%s.%s", registerRequest.Name, actionName)
		logger.Info("registering command", "name", name)

		applicationCommand := action.ToApplicationCommand(registerRequest.Name)

		if err := pluginHost.brain.CreateCommand(ctx, applicationCommand); err != nil {
			return fmt.Errorf("create command '%s': %w", name, err)
		}
	}

	return nil
}

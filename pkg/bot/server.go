package bot

import (
	"context"
	"errors"

	pluginostv1alpha1 "github.com/bryanl/lilbot/gen/go/pluginhost/v1alpha1"
)

// Server is the plugin host server.
type Server struct {
	pluginostv1alpha1.UnsafeLilbotPluginHostServiceServer

	pluginHost PluginHost
}

// NewServer creates an instance of Server.
func NewServer(pluginHost PluginHost) *Server {
	server := &Server{
		pluginHost: pluginHost,
	}

	return server
}

// Register registers a plugin with the plugin host.
func (s Server) Register(ctx context.Context, request *pluginostv1alpha1.RegisterRequest) (*pluginostv1alpha1.RegisterResponse, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}

	resp := &pluginostv1alpha1.RegisterResponse{}

	return resp, nil
}

package bot

import (
	"context"
	"errors"
	"fmt"
	"net"

	pluginostv1alpha1 "github.com/bryanl/lilbot/gen/go/pluginhost/v1alpha1"
	"github.com/bryanl/lilutil/log"
	"google.golang.org/grpc"
)

// GRPCServer create a GRPC server.
type GRPCServer struct {
	listener   net.Listener
	pluginHost PluginHost
}

// NewGRPCServer creates an instance of GRPCServer.
func NewGRPCServer(addr string, pluginHost PluginHost) (*GRPCServer, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("create listener: %w", err)
	}

	if pluginHost == nil {
		return nil, errors.New("plugin host is nil")
	}

	server := &GRPCServer{
		listener:   l,
		pluginHost: pluginHost,
	}

	return server, nil
}

// Start starts the GRPC server.
func (server *GRPCServer) Start(ctx context.Context) (<-chan struct{}, error) {
	logger := log.From(ctx).WithName("lilbot.GRPCServer")

	s := grpc.NewServer()

	pluginostv1alpha1.RegisterLilbotPluginHostServiceServer(s, NewServer(server.pluginHost))

	logger.Info("starting GRPC server", "addr", server.listener.Addr().String())

	go func() {
		if err := s.Serve(server.listener); err != nil {
			logger.Error(err, "stop server")
		}

		logger.Info("server has stopped")
	}()

	ch := make(chan struct{}, 1)

	go func() {
		<-ctx.Done()
		logger.Info("stopping gracefully")
		s.GracefulStop()
		close(ch)
	}()

	return ch, nil
}

package bot

// PluginHost is a plugin host interface.
type PluginHost interface {
	Register() error
}

// LivePluginHost is a live implementation of PluginHost.
type LivePluginHost struct{}

var _ PluginHost = &LivePluginHost{}

// NewPluginHost creates an instance of LivePluginHost.
func NewPluginHost() *LivePluginHost {
	pluginHost := &LivePluginHost{}

	return pluginHost
}

// Register registers a plugin with the host.
func (pluginHost *LivePluginHost) Register() error {
	panic("implement me")
}

package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)


var (
	// HandshakeConfig is the config for the handshake between lpass and a
	// plugin. It is provided here as an easy way for everyone to share the same config.
	HandshakeConfig plugin.HandshakeConfig = goplugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "LPASS_MAGIC_COOKIE",
		MagicCookieValue: "2",
	}
)

// LastPass defines the interface a lastpass plugin must fulfill.
type LastPass interface {
	Execute(args []string)
}

// LastPassPlugin ...
type LastPassPlugin struct {
	Impl LastPass
}

// LastPassRPCServer ...
type LastPassRPCServer struct {
	Impl LastPass
}

// LastPassRPC is an implementation of `LastPass` that talks over RPC.
type LastPassRPC struct {
	client *rpc.Client
}

// Server ...
func (p *LastPassPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &LastPassRPCServer{Impl: p.Impl}, nil
}

// Client ...
func (LastPassPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &LastPassRPC{client: c}, nil
}

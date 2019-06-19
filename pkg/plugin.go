package plugin

import (
	"net/rpc"
	"github.com/hashicorp/go-plugin"
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
	return &LastPassRPCServer{Impl: p.Impl }, nil
}

// Client ...
func (LastPassPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &LastPassRPC{client: c}, nil
}

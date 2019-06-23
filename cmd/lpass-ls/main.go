package main

import (
	"github.com/hashicorp/go-plugin"
	lpassPlugin "github.com/nick96/lastpass-go/pkg/plugin"
	log "github.com/sirupsen/logrus"
)

const (
	// PluginName is the name of this plugin
	PluginName string = "lpass-ls"
)

// LpassLsPlugin is an implementation of the LastPass interface.
type LpassLsPlugin struct{}

// Execute executes the functionality of the `lpass-ls` plugin. That is, it
// lists the available credentials stored int he lastpass account in a manner
// similar to that done by the `ls` *nix command.
func (p *LpassLsPlugin) Execute(args []string) {
	log.Debugf("Executing %s with arguments: %v", PluginName, args)
}

func main() {
	lsPlugin := &LpassLsPlugin{}
	pluginMap := map[string]plugin.Plugin{
		"ls": &lpassPlugin.LastPassPlugin{
			Impl: lsPlugin,
		},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: lpassPlugin.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

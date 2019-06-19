package main

import (
	"os/exec"

	goplugin "github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"

	"os"

	lpassPlugin "github.com/nick96/lastpass-go/pkg"
)

const (
	// PluginPrefix is the prefix to look for in plugins
	PluginPrefix string = "lpass"
)

func main() {
	pluginPaths := []string{}
	pluginMap, err := lpassPlugin.Map(PluginPrefix, pluginPaths)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Plugin map: %v", pluginMap)

	handshakeConfig := goplugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}

	plugin := os.Args[1]
	pluginPath, err := lpassPlugin.ExpandName(plugin, PluginPrefix, pluginPaths)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Expanded plugin %s to path %s", plugin, pluginPath)

	client := goplugin.NewClient(&goplugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(pluginPath),
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal("Error: Could not get client: %v", err)
	}

	raw, err := rpcClient.Dispense(os.Args[1])
	if err != nil {
		log.Fatal("Error: Could not dispense greeter: %v", err)
	}

	lastpass := raw.(lpassPlugin.LastPass)
	lastpass.Execute(os.Args[2:])
}


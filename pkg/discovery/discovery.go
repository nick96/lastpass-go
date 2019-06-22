package discovery

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	goplugin "github.com/hashicorp/go-plugin"

	lpassPlugin "github.com/nick96/lastpass-go/pkg/plugin"
)

type listDir func(string) ([]os.FileInfo, error)

//nolint:gochecknoglobals
var (
	// Implementation of the listDir type that works on the actual file system
	listDirFs listDir = ioutil.ReadDir
)

// PluginMap returns a map of plugin names to the corresponding plugin object.
func PluginMap(prefix string, pluginPaths []string) (map[string]goplugin.Plugin, error) {
	return pluginMap(prefix, pluginPaths, listDirFs)
}

// ExpandName expands the name of plugin to the path to its corresponding executable.
func ExpandName(name, prefix string, pluginPaths []string) (string, error) {
	return expandName(name, prefix, pluginPaths, listDirFs)
}

func pluginMap(prefix string, pluginPaths []string, listDir listDir) (map[string]goplugin.Plugin, error) {
	pluginMap := map[string]goplugin.Plugin{}
	plugins, err := findPlugins(prefix, pluginPaths, listDir)
	if err != nil {
		return pluginMap, err
	}

	for _, plugin := range plugins {
		name := strings.TrimPrefix(path.Base(plugin), prefix)
		pluginMap[name] = &lpassPlugin.LastPassPlugin{}
	}

	return pluginMap, nil
}

func expandName(name, prefix string, pluginPaths []string, listDir listDir) (string, error) {
	plugins, err := findPlugins(prefix, pluginPaths, listDir)
	if err != nil {
		return "", fmt.Errorf("could not expand plugin %s: %v", name, err)
	}

	for _, plugin := range plugins {
		if strings.HasSuffix(plugin, name) {
			return plugin, nil
		}
	}
	return "", fmt.Errorf("could not find plugin %s", name)
}

// Find all available plugins and return the absolute path to them.
func findPlugins(prefix string, pluginPaths []string, listDir listDir) ([]string, error) {
	// Get the plugins in the path
	plugins, err := findPluginsInPath(prefix, listDir)
	if err != nil {
		return []string{}, err
	}

	for _, pluginPath := range pluginPaths {
		foundPlugins, err := findPluginsInDirectory(prefix, pluginPath, listDir)
		if err != nil {
			return []string{}, err
		}
		plugins = append(plugins, foundPlugins...)
	}
	return plugins, nil
}

// Find all plugins in the path
func findPluginsInPath(prefix string, listDir listDir) ([]string, error) {
	path := os.Getenv("PATH")
	plugins := make([]string, 0)
	for _, dir := range filepath.SplitList(path) {
		dirPlugins, err := findPluginsInDirectory(prefix, dir, listDir)
		if err != nil {
			return []string{}, err
		}
		plugins = append(plugins, dirPlugins...)
	}
	return plugins, nil
}

// findPluginsInDirectory recursively finds plugins with a given prefix in a directory.
func findPluginsInDirectory(prefix, dir string, listDir listDir) ([]string, error) {
	plugins := []string{}
	files, err := listDir(dir)
	if err != nil {
		return []string{}, err
	}

	for _, file := range files {
		if file.IsDir() {
			belowPlugins, err := findPluginsInDirectory(prefix, file.Name(), listDir)
			if err != nil {
				return []string{}, err
			}
			plugins = append(plugins, belowPlugins...)
		} else if strings.HasPrefix(file.Name(), prefix) {
			plugins = append(plugins, file.Name())
		}
	}
	return plugins, nil
}

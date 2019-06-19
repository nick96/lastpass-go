package plugin

import (
	goplugin "github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"fmt"
)


// Map returns a map of plugin names to the corresponding plugin object.
func Map(prefix string, pluginPaths []string) (map[string]goplugin.Plugin, error) {
	pluginMap := map[string]goplugin.Plugin{}
	plugins, err := findPlugins(prefix, pluginPaths)
	if err != nil {
		return pluginMap, err
	}

	for _, plugin := range plugins {
		name := strings.TrimPrefix(path.Base(plugin), prefix)
		pluginMap[name] = &LastPassPlugin{}
	}

	return pluginMap, nil
}

// ExpandName expands the name of plugin to the path to its corresponding executable.
func ExpandName(name, prefix string, pluginPaths []string) (string, error) {
	plugins, err := findPlugins(prefix, pluginPaths)
	if err != nil {
		return "", fmt.Errorf("Error: Could not expand plugin %s: %v", name, err)
	}

	for _, plugin := range plugins {
		if strings.HasSuffix(plugin, name) {
			return plugin, nil
		}
	}
	return "", fmt.Errorf("Error: Could not find plugin %s", name)
}

// Find all available plugins and return the absolute path to them.
func findPlugins(prefix string, pluginPaths []string) ([]string, error) {
	// Get the plugins in the path
	plugins, err := findPluginsInPath(prefix)
	if err != nil {
		return []string{}, err
	}

	for _, pluginPath := range pluginPaths {
		foundPlugins, err := findPluginsInDirectory(prefix, pluginPath)
		if err != nil {
			return []string{}, err
		}
		plugins = append(plugins, foundPlugins...)
	}
	return plugins, nil
}

// Find all plugins in the path
func findPluginsInPath(prefix string) ([]string, error) {
	path := os.Getenv("PATH")
	plugins := make([]string, 0)
	for _, dir := range path.SplitList(path) {
		dirPlugins, err := findPluginsInDirectory(prefix, dir)
		if err != nil {
			return []string{}, err
		}
		plugins := append(plugins, dirPlugins...)
	}
	return plugins
}

// findPluginsInDirectory recursively finds plugins with a given prefix in a directory.
func findPluginsInDirectory(prefix, dir string) ([]string, error) {
	plugins := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	for _, file := range files {
		if (file.IsDir()) {
			belowPlugins, err := findPluginsInDirectory(prefix, file.Name())
			if err != nil {
				return []string{}, err
			}
			plugins = append(plugins, belowPlugins)
		} else if strings.HasPrefix(file.Name(), prefix) {
			plugins = append(plugins, file.Name())
		}
	}
	return plugins
}

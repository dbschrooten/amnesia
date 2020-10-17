package plugin

import (
	"amnesia/src/config"
	"amnesia/src/service"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

var (
	LoadedPlugins = make(map[string]service.ServiceImpl)
)

func mapPlugin(path string) error {
	plug, err := plugin.Open(path)

	if err != nil {
		return err
	}

	impl, err := plug.Lookup("Implementation")

	if err != nil {
		return err
	}

	implPlug, ok := impl.(service.ServiceImpl)

	if !ok {
		return fmt.Errorf("Unexpected type from module symbol")
	}

	var pluginInfo = implPlug.Info()

	// validate and sanitize plugin name, before loading
	LoadedPlugins[pluginInfo["name"]] = implPlug
	service.PluginImplementations = append(service.PluginImplementations, pluginInfo["name"])

	log.Printf("Loaded Plugin: %s", implPlug.Info()["name"])

	return nil
}

func loader() error {
	if err := filepath.Walk(config.PluginFolder, func(path string, info os.FileInfo, err error) error {
		// only scan for .so plugin files
		if !info.IsDir() {
			if filepath.Ext(path) == ".so" {
				mapPlugin(path)
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func Setup() error {
	if err := loader(); err != nil {
		return err
	}

	return nil
}

package plugin

import (
	"log"
)

var plugins []*Plugin

func init() {
	log.Println("Plugins start")
}

type Plugin struct {
	Name string
}

func Register(p *Plugin) {
	plugins = append(plugins, p)
}

func ListPlugins() []*Plugin {
	return plugins
}

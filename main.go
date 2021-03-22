package main

import (
	"log"
	"os"
	"plugin"

	myPlugin "github.com/nyanpassu/go-plugin-main/plugin"
)

func main() {
	pluginPath := os.Args[1]
	log.Println("Loading: " + pluginPath)

	_, err := plugin.Open(pluginPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range myPlugin.ListPlugins() {
		log.Println(p.Name)
	}
}

package main

import (
	"github.com/Israel-Ferreira/transform-response-plugin/pkg/plugins"
	"github.com/Kong/go-pdk/server"
)

var pluginVersion = "0.0.2"
var priority = 1

func main() {
	server.StartServer(plugins.New, pluginVersion, priority)
}

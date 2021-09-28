package main

import (
	"github.com/judell/steampipe-plugin-hypothesis/hypothesis"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hypothesis.Plugin})
}

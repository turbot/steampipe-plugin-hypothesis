package main

import (
	"github.com/turbot/steampipe-plugin-hypothesis/hypothesis"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hypothesis.Plugin})
}

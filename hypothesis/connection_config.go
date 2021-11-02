package hypothesis

import (
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type hypothesisConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &hypothesisConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hypothesisConfig {
	if connection == nil || connection.Config == nil {
		fmt.Printf("%v+\n", connection)
		return hypothesisConfig{}
	}
	config, _ := connection.Config.(hypothesisConfig)
	fmt.Printf("%v+\n", config)
	return config
}

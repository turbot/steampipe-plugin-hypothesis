package hypothesis

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
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
		return hypothesisConfig{}
	}
	config, _ := connection.Config.(hypothesisConfig)
	return config
}

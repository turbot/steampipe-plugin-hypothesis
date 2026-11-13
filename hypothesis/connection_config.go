package hypothesis

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type hypothesisConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &hypothesisConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hypothesisConfig {
	if connection == nil || connection.GetConfig() == nil {
		return hypothesisConfig{}
	}
	config, _ := connection.GetConfig().(hypothesisConfig)
	return config
}

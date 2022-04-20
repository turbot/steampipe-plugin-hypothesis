package hypothesis

import (
	"context"
	//"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-hypothesis",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"hypothesis_search": tableHypothesisSearch(ctx),
			"hypothesis_profile": tableHypothesisProfile(ctx),
		},
	}
	return p
}

package hypothesis

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-hypothesis",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"hypothesis_search": tableHypothesisSearch(ctx),
		},
	}
	return p
}

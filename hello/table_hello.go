package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

type Hello struct {
	ID       int    `json:"id"`
	Greeting string `json:"greeting"`
	JSON	 string `json:"json"`
}

func tableHello(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello",
		Description: "Simplest Steampipe plugin",
		List: &plugin.ListConfig{
			Hydrate: listGreeting,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "an int"},
			{Name: "greeting", Type: proto.ColumnType_STRING, Description: "a string"},
			{Name: "json", Type: proto.ColumnType_JSON, Description: "a json object"},
		},
	}
}

func listGreeting(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 1; i <= 3; i++ {
		plugin.Logger(ctx).Info("listGreeting", "number", i)		
		greeting := Hello{i, "Hello", "{\"hello\": \"world\"}"}
		d.StreamListItem(ctx, &greeting)
	}
	return nil, nil
}

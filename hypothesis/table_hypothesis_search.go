package hypothesis

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
    hyp "github.com/judell/hypothesis-go"
)

func tableHypothesisSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hypothesis_search",
		Description: "Search for Hypothesis annotations",
		List: &plugin.ListConfig{
			Hydrate: listSearchResults,
			KeyColumns: plugin.SingleColumn("query"),			
		},
		Columns: []*plugin.Column{
			{Name: "query", Type: proto.ColumnType_STRING, Hydrate: queryString, Transform: transform.FromValue(), Description: "The search query."},
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "created", Type: proto.ColumnType_STRING},
			{Name: "updated", Type: proto.ColumnType_STRING},
			{Name: "user", Type: proto.ColumnType_STRING},
			{Name: "group", Type: proto.ColumnType_STRING},
			{Name: "uri", Type: proto.ColumnType_STRING},
			{Name: "text", Type: proto.ColumnType_STRING},
			{Name: "tags", Type: proto.ColumnType_JSON},
			{Name: "document", Type: proto.ColumnType_JSON},
			{Name: "target", Type: proto.ColumnType_JSON},
		},
	}
}

func listSearchResults(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	qs := quals["query"].GetStringValue()

	plugin.Logger(ctx).Warn("hypothesis.listSearchResults", "qs", qs)

	hypothesisConfig := GetConfig(d.Connection)

	token := os.Getenv("H_TOKEN")
    if hypothesisConfig.Token != nil {
		token = *hypothesisConfig.Token
	}

	u, err := url.Parse("https://hypothes.is/api/search?" + qs)
	if err != nil {
		panic(err)
		}
	
    m, _ := url.ParseQuery(u.RawQuery)
	plugin.Logger(ctx).Warn("hypothesis.listSearchResults", "m", fmt.Sprintf("%+v", m))

	searchParams := hyp.SearchParams{}
	if mapContainsKey(m, "any") {
		searchParams.Any = m["any"][0]
	}
	if mapContainsKey(m, "user") {
		searchParams.User = m["user"][0]
	}
	if mapContainsKey(m, "group") {
		searchParams.Group = m["group"][0]
	}
	if mapContainsKey(m, "uri") {
		searchParams.Uri = m["uri"][0]
	} else {
		if mapContainsKey(m, "wildcard_uri") {
			searchParams.WildcardUri = m["wildcard_uri"][0]
		}
	}
	plugin.Logger(ctx).Warn("hypothesis.listSearchResults", "searchParams", fmt.Sprintf("%+v", searchParams))
	    
	client := hyp.NewClient(
		token,
		searchParams,
		0, // use library default: 2000
	)

	plugin.Logger(ctx).Warn("hypothesis.listSearchResults", "client", fmt.Sprintf("%+v", client))

	rows, err := client.SearchAll()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}

func queryString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	return q, nil
}

func mapContainsKey(m map[string][]string, key string) bool {
	if len(m[key]) > 0 && m[key][0] != "" {
		return true
	}	
	return false
}

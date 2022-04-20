package hypothesis

import (
	"context"
	"fmt"
	"net/url"
	"os"

	hyp "github.com/judell/hypothesis-go"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableHypothesisSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hypothesis_search",
		Description: "Search for Hypothesis annotations",
		List: &plugin.ListConfig{
			Hydrate:    listSearchResults,
			KeyColumns: plugin.SingleColumn("query"),
		},
		Columns: []*plugin.Column{
			{Name: "query", Type: proto.ColumnType_STRING, Hydrate: queryString, Transform: transform.FromValue(), Description: "The search query."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The annotation id, works with https://hypothes.is/a/{ID}."},
			{Name: "created", Type: proto.ColumnType_STRING, Description: "The creation date of the annotation."},
			{Name: "updated", Type: proto.ColumnType_STRING, Description: "The last update date of the annotation."},
			{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User").Transform(userIdToUsername), Description: "The Hypothesis username of the person who created the annotation."},
			{Name: "group_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group"), Description: "The annotation's group: __world__ or a private group id."},
			{Name: "uri", Type: proto.ColumnType_STRING, Description: "URL of the annotated resource."},
			{Name: "text", Type: proto.ColumnType_STRING, Description: "Textual body of the annotation, as MarkDown/HTML."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags on the annotation, as a JSONB array of strings."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Document").Transform(documentToTitle), Description: "The HTML doctitle of the annotated URL."},
			{Name: "document", Type: proto.ColumnType_JSON, Description: "An element that contains the title and maybe other metadata."},
			{Name: "target", Type: proto.ColumnType_JSON, Description: "The selectors that define the document selection to which the annotation anchors."},
			{Name: "exact", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target").Transform(selectorsToExact), Description: "The text of the selection (aka quote) to which the annotation anchors."},
			{Name: "refs", Type: proto.ColumnType_JSON, Transform: transform.FromField("References"), Description: "IDs forming the reference chain to which this annotation belongs."},

		},
	}
}

func listSearchResults(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	qs := quals["query"].GetStringValue()

	plugin.Logger(ctx).Info("hypothesis.listSearchResults", "qs", qs)

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
	plugin.Logger(ctx).Info("hypothesis.listSearchResults", "m", fmt.Sprintf("%+v", m))

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
	} else if mapContainsKey(m, "wildcard_uri") {
		searchParams.WildcardUri = m["wildcard_uri"][0]
	}
	if mapContainsKey(m, "tag") {
		searchParams.Tags = append(searchParams.Tags, m["tag"]...)
	}
	if mapContainsKey(m, "limit") {
		searchParams.Limit = m["limit"][0]
	}

	plugin.Logger(ctx).Info("hypothesis.listSearchResults", "searchParams", fmt.Sprintf("%+v", searchParams))

	client := hyp.NewClient(
		token,
		searchParams,
		0, // use library default
	)

	plugin.Logger(ctx).Info("hypothesis.listSearchResults", "client", fmt.Sprintf("%+v", client))

	i := 0
	for row := range client.SearchAll() {
		i += 1
		if i % 500 == 0 {
			plugin.Logger(ctx).Info("hypothesis.listSearchResults", "row", fmt.Sprintf(`%d`, i))
		}

		d.StreamListItem(ctx, row)

		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			break
		}

	}
	return nil, nil
}

package hypothesis

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"regexp"

	hyp "github.com/judell/hypothesis-go"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The annotation id, works with https://hypothes.is/a/{ID}"},
			{Name: "created", Type: proto.ColumnType_STRING, Description: "The creation date of the annotation"},
			{Name: "updated", Type: proto.ColumnType_STRING, Description: "The last update date of the annotation"},
			{Name: "user", Type: proto.ColumnType_STRING, Transform: transform.FromField("User").Transform(userIdToUsername), Description: "The Hypothesis username of the person who created the annotation"},
			{Name: "group", Type: proto.ColumnType_STRING, Description: "The annotation's group: __world__ or a private group id"},
			{Name: "uri", Type: proto.ColumnType_STRING, Description: "URL of the annotated resource"},
			{Name: "text", Type: proto.ColumnType_STRING, Description: "Textual body of the annotation, as MarkDown/HTML"},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags on the annotation, as a JSONB array of strings"},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Document").Transform(documentToTitle), Description: "The HTML doctitle of the annotated URL."},
			{Name: "document", Type: proto.ColumnType_JSON, Description: "An element that contains the title and maybe other metadata"},
			{Name: "target", Type: proto.ColumnType_JSON, Description: "The selectors that define the document selection to which the annotation anchors"},
			{Name: "exact", Type: proto.ColumnType_STRING, Transform: transform.FromField("Target").Transform(targetToExact), Description: "The text of the selection (aka quote) to which the annotation anchors"},
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
	} else if mapContainsKey(m, "wildcard_uri") {
		searchParams.WildcardUri = m["wildcard_uri"][0]
	}
	if mapContainsKey(m, "tag") {
		searchParams.Tags = append(searchParams.Tags, m["tag"]...)
	}

	plugin.Logger(ctx).Warn("hypothesis.listSearchResults", "searchParams", fmt.Sprintf("%+v", searchParams))

	client := hyp.NewClient(
		token,
		searchParams,
		0, // use library default
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

func documentToTitle(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	doc := input.Value.(struct {
		Title []string "json:\"title\""
	})
	return doc.Title[0], nil
}

func userIdToUsername(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	userId := input.Value.(string)
	re := regexp.MustCompile("acct:|@hypothes.is")
	userName := re.ReplaceAllString(userId, "")
	return userName, nil
}


type Target = []struct {
	Source   string `json:"source"`
	Selector []struct {
		End    int    `json:"end,omitempty"`
		Type   string `json:"type"`
		Start  int    `json:"start,omitempty"`
		Exact  string `json:"exact,omitempty"`
		Prefix string `json:"prefix,omitempty"`
		Suffix string `json:"suffix,omitempty"`
	} `json:"selector"`
}

func targetToExact(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	empty := ""
	target := input.Value.(Target)
	if len(target) == 0 {
		return empty, nil
	}
	if len(target[0].Selector) == 0 {
		return empty, nil
	}
	for _, sel := range target[0].Selector {
		if sel.Type == "TextQuoteSelector" {
			return sel.Exact, nil
		}
	}
	return empty, nil
}

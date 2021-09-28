package hypothesis

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"	

)

type SearchResult struct {
	Total int `json:"total"`
	Rows []Row `json:"rows"`
  }
  
type Row struct {
	ID          string        `json:"id"`
	Created     time.Time     `json:"created"`
	Updated     time.Time     `json:"updated"`
	User        string        `json:"user"`
	URI         string        `json:"uri"`
	Text        string        `json:"text"`
	Tags        []string      `json:"tags"`
	Group       string        `json:"group"`
	Target []struct {
		Source   string `json:"source"`
		Selector []struct {
			End    int    `json:"end,omitempty"`
			Type   string `json:"type"`
			Start  int    `json:"start,omitempty"`
			Exact  string `json:"exact,omitempty"`
			Prefix string `json:"prefix,omitempty"`
			Suffix string `json:"suffix,omitempty"`
		} `json:"selector"`
	} `json:"target"`
	Document struct {
		Title []string `json:"title"`
	} `json:"document"`
	UserInfo struct {
		DisplayName string `json:"display_name"`
	} `json:"user_info"`
}

func tableHypothesisSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hypothesis_search",
		Description: "Search for Hypothesis annotations",
		List: &plugin.ListConfig{
			Hydrate: listSearchResults,
			//KeyColumns: plugin.SingleColumn("query"),			
		},
		Columns: []*plugin.Column{
//			{Name: "query", Type: proto.ColumnType_STRING},
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "created", Type: proto.ColumnType_STRING},
			{Name: "userid", Type: proto.ColumnType_STRING, Transform: transform.FromField("User")}, 
			{Name: "groupid", Type: proto.ColumnType_STRING, Transform: transform.FromField("Group")},
			{Name: "uri", Type: proto.ColumnType_STRING},
			{Name: "text", Type: proto.ColumnType_STRING},
			{Name: "tags", Type: proto.ColumnType_JSON},
			{Name: "target", Type: proto.ColumnType_JSON},
			{Name: "document", Type: proto.ColumnType_JSON},
			{Name: "user_info", Type: proto.ColumnType_JSON},
		},
	}
}

func listSearchResults(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var client http.Client
	url := "https://hypothes.is/api/search"
	r, err := client.Get(url)
	if err != nil {
		plugin.Logger(ctx).Error("Error getting Hypothesis search results", "", err)
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var searchResult SearchResult
	_ = decoder.Decode(&searchResult)
	plugin.Logger(ctx).Error("INFO", "", searchResult.Total)
	for _, row := range searchResult.Rows {
		d.StreamListItem(ctx, row)
	}
	return nil, nil
}

func queryString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	return q, nil
}
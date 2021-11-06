package hypothesis

import (
	"context"
	"os"

	hyp "github.com/judell/hypothesis-go"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableHypothesisProfile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hypothesis_profile",
		Description: "Profile for logged-in Hypothesis user",
		List: &plugin.ListConfig{
			Hydrate: listProfile,
		},
		Columns: []*plugin.Column{
			{Name: "user", Type: proto.ColumnType_STRING, Transform: transform.FromField("Userid").Transform(userIdToUsername), Description: "The user whose profile to get."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("UserInfo").Transform(userInfoToDisplayName), Description: "The user's display name (if any)."},
			{Name: "authority", Type: proto.ColumnType_STRING, Description: "Authority of the user's account."},
			{Name: "groups", Type: proto.ColumnType_JSON, Description: "The user's groups"},
		},
	}
}

func listProfile(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	hypothesisConfig := GetConfig(d.Connection)

	token := os.Getenv("H_TOKEN")
	if hypothesisConfig.Token != nil {
		token = *hypothesisConfig.Token
	}

	if token == "" {
		return hyp.Profile{}, nil
	}

	client := hyp.NewClient(
		token,
		hyp.SearchParams{},
		0, // use library default
	)

	profile, err := client.Profile()
	if err != nil {
		panic(err)
	}

	d.StreamListItem(ctx, profile)
	return nil, nil
}

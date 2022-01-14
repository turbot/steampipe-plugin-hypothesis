package hypothesis

import (
	"regexp"
	"context"

	hyp "github.com/judell/hypothesis-go"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

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
	if len(doc.Title) == 0 {
		return "untitled",  nil
	}
	return doc.Title[0], nil
}

func selectorsToExact(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	targets := input.Value.([]hyp.Target)
	selectors := targets[0].Selector
	exact, err := hyp.SelectorsToExact(selectors)
	if err != nil {
		return "", nil
	}
	return exact, nil
}

func userIdToUsername(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	userId := input.Value.(string)
	re := regexp.MustCompile("acct:|@hypothes.is")
	userName := re.ReplaceAllString(userId, "")
	return userName, nil
}

func userInfoToDisplayName(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	userInfo := input.Value.(struct {DisplayName string `json:"display_name"`})
	return userInfo.DisplayName, nil
}


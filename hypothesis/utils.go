package hypothesis

import(
	"regexp"
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

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


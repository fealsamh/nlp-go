package aibot

import "net/http"

type deployRequest struct {
	Bot *Bot `json:"bot"`
}

// DeployResult ...
type DeployResult string

var results = map[int]DeployResult{
	http.StatusCreated:  "created",
	http.StatusAccepted: "updated",
}

// Deploy ...
func (c *Client) Deploy(bot *Bot) (DeployResult, error) {
	s, err := c.callServicePost("/api/v1/bots", &deployRequest{
		Bot: bot,
	}, nil)
	if err != nil {
		return "", err
	}
	return results[s], nil
}

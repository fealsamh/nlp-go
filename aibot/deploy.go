package aibot

import "net/http"

type DeployRequest struct {
	Bot *Bot `json:"bot"`
}

type DeployResult string

var results = map[int]DeployResult{
	http.StatusCreated:  "created",
	http.StatusAccepted: "updated",
}

func (c *Client) Deploy(bot *Bot) (DeployResult, error) {
	s, err := c.callServicePost("/api/v1/bots", &DeployRequest{
		Bot: bot,
	}, nil)
	if err != nil {
		return "", err
	}
	return results[s], nil
}

package aibot

type ListBotsResponse struct {
	BotIDs []string `json:"bot_ids"`
}

func (c *Client) ListBots() ([]string, error) {
	var r *ListBotsResponse
	if _, err := c.callServiceGet("/api/v1/bots", &r); err != nil {
		return nil, err
	}
	return r.BotIDs, nil
}

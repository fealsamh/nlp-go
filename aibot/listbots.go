package aibot

type listBotsResponse struct {
	BotIDs []string `json:"bot_ids"`
}

// ListBots ...
func (c *Client) ListBots() ([]string, error) {
	var r *listBotsResponse
	if _, err := c.callServiceGet("/api/v1/bots", &r); err != nil {
		return nil, err
	}
	return r.BotIDs, nil
}

package aibot

type getBotResponse struct {
	Bot *Bot `json:"bot"`
}

// GetBot ...
func (c *Client) GetBot(id string) (*Bot, error) {
	var r *getBotResponse
	if _, err := c.callServiceGet("/api/v1/bots/"+id, &r); err != nil {
		return nil, err
	}
	return r.Bot, nil
}

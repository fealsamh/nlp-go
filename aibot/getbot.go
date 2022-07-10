package aibot

type GetBotResponse struct {
	Bot *Bot `json:"bot"`
}

func (c *Client) GetBot(id string) (*Bot, error) {
	var r *GetBotResponse
	if _, err := c.callServiceGet("/api/v1/bots/"+id, &r); err != nil {
		return nil, err
	}
	return r.Bot, nil
}

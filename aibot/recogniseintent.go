package aibot

type recogniseIntentRequest struct {
	BotID string `json:"bot_id"`
	Text  string `json:"text"`
}

type IntentSimilarity struct {
	IntentID   string  `json:"intent_id"`
	Similarity float64 `json:"similarity"`
}

type recogniseIntentResponse struct {
	Intents  []IntentSimilarity  `json:"intents"`
	Entities map[string][]string `json:"entities"`
}

func (c *Client) RecogniseIntent(botId, text string) ([]IntentSimilarity, map[string][]string, error) {
	var r *recogniseIntentResponse
	if _, err := c.callServicePost("/api/v1/bots/intents", &recogniseIntentRequest{
		BotID: botId,
		Text:  text,
	}, &r); err != nil {
		return nil, nil, err
	}
	return r.Intents, r.Entities, nil
}

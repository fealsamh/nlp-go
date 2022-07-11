package aibot

type RecogniseIntentRequest struct {
	BotID string `json:"bot_id"`
	Text  string `json:"text"`
}

type IntentSimilarity struct {
	IntentID   string  `json:"intent_id"`
	Similarity float64 `json:"similarity"`
}

type RecogniseIntentResponse struct {
	Intents []IntentSimilarity `json:"intents"`
}

func (c *Client) RecogniseIntent(botId, text string) ([]IntentSimilarity, error) {
	var r *RecogniseIntentResponse
	if _, err := c.callServicePost("/api/v1/bots/intents", &RecogniseIntentRequest{
		BotID: botId,
		Text:  text,
	}, &r); err != nil {
		return nil, err
	}
	return r.Intents, nil
}

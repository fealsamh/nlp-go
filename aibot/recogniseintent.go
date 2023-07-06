package aibot

type recogniseIntentRequest struct {
	BotID string `json:"bot_id"`
	Text  string `json:"text"`
}

// IntentSimilarity ...
type IntentSimilarity struct {
	IntentID   string  `json:"intent_id"`
	Similarity float64 `json:"similarity"`
}

type recogniseIntentResponse struct {
	Intents  []IntentSimilarity  `json:"intents"`
	Entities map[string][]string `json:"entities"`
}

// RecogniseIntent ...
func (c *Client) RecogniseIntent(botID, text string) ([]IntentSimilarity, map[string][]string, error) {
	var r *recogniseIntentResponse
	if _, err := c.callServicePost("/api/v1/bots/intents", &recogniseIntentRequest{
		BotID: botID,
		Text:  text,
	}, &r); err != nil {
		return nil, nil, err
	}
	return r.Intents, r.Entities, nil
}

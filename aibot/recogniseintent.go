package aibot

import "encoding/json"

type RecogniseIntentRequest struct {
	BotID string `json:"bot_id"`
	Text  string `json:"text"`
}

type RecogniseIntentResponse struct {
	Chart json.RawMessage `json:"chart"`
}

func (c *Client) RecogniseIntent(botId, text string) (string, error) {
	var r *RecogniseIntentResponse
	if _, err := c.callServicePost("/api/v1/bots/intents", &RecogniseIntentRequest{
		BotID: botId,
		Text:  text,
	}, &r); err != nil {
		return "", err
	}
	return string(r.Chart), nil
}

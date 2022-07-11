package main

import (
	"fmt"

	"github.com/fealsamh/nlp-go/aibot"
)

func recogniseIntent(cl *aibot.Client, botId, text string) {
	r, err := cl.RecogniseIntent(botId, text)
	if err != nil {
		exitWithError(err)
	}
	for _, is := range r {
		fmt.Printf("%s: %f\n", is.IntentID, is.Similarity)
	}
}

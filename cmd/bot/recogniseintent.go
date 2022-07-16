package main

import (
	"fmt"
	"strings"

	"github.com/fealsamh/nlp-go/aibot"
)

func recogniseIntent(cl *aibot.Client, botId, text string) {
	r, e, err := cl.RecogniseIntent(botId, text)
	if err != nil {
		exitWithError(err)
	}
	fmt.Println("intents:")
	for _, is := range r {
		fmt.Printf("  %s: %f\n", is.IntentID, is.Similarity)
	}
	if len(e) > 0 {
		fmt.Println("entities:")
		for k, v := range e {
			fmt.Printf("  %s: %s\n", k, strings.Join(v, ", "))
		}
	}
}

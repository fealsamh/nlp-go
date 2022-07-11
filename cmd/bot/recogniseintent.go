package main

import (
	"fmt"

	"github.com/fealsamh/nlp-go/aibot"
)

func recogniseIntent(cl *aibot.Client, botId, text string) {
	s, err := cl.RecogniseIntent(botId, text)
	if err != nil {
		exitWithError(err)
	}
	fmt.Println(s)
}

package main

import (
	"fmt"

	"github.com/fealsamh/nlp-go/aibot"
)

func listBots(cl *aibot.Client) {
	l, err := cl.ListBots()
	if err != nil {
		exitWithError(err)
	}
	for _, b := range l {
		fmt.Println(b)
	}
}

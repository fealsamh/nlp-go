package main

import (
	"fmt"
	"regexp"

	"github.com/fealsamh/nlp-go/aibot"
)

func listBots(cl *aibot.Client, re *regexp.Regexp) {
	l, err := cl.ListBots()
	if err != nil {
		exitWithError(err)
	}
	for _, id := range l {
		if re == nil || re.MatchString(id) {
			fmt.Println(id)
		}
	}
}

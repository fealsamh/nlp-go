package main

import (
	"os"

	"github.com/fealsamh/nlp-go/aibot"
	"gopkg.in/yaml.v3"
)

func getBot(cl *aibot.Client, id string) {
	b, err := cl.GetBot(id)
	if err != nil {
		exitWithError(err)
	}
	if err := yaml.NewEncoder(os.Stdout).Encode(b); err != nil {
		exitWithError(err)
	}
}

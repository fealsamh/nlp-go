package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fealsamh/nlp-go/aibot"
	"gopkg.in/yaml.v3"
)

func deploy(cl *aibot.Client, fp string) {
	f, err := os.Open(fp)
	if err != nil {
		exitWithError(err)
	}
	defer f.Close()
	deployReader(cl, f)
}

func deployReader(cl *aibot.Client, r io.Reader) {
	var bot *aibot.Bot
	if err := yaml.NewDecoder(r).Decode(&bot); err != nil {
		exitWithError(err)
	}

	if err := bot.Validate(); err != nil {
		exitWithError(err)
	}
	fmt.Println("the bot is valid")

	dr, err := cl.Deploy(bot)
	if err != nil {
		exitWithError(err)
	}

	fmt.Println(dr)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fealsamh/nlp-go/aibot"
	"gopkg.in/yaml.v3"
)

func check(fp string, printJSON bool) {
	f, err := os.Open(fp)
	if err != nil {
		exitWithError(err)
	}
	defer f.Close()
	checkReader(f, printJSON)
}

func checkReader(r io.Reader, printJSON bool) {
	var bot *aibot.Bot
	if err := yaml.NewDecoder(r).Decode(&bot); err != nil {
		exitWithError(err)
	}

	if err := bot.Validate(); err != nil {
		exitWithError(err)
	}
	fmt.Println("the bot is valid")

	if printJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		if err := enc.Encode(bot); err != nil {
			exitWithError(err)
		}
	}
}

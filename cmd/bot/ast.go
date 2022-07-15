package main

import (
	"fmt"
	"strings"
	"text/scanner"

	"github.com/fealsamh/nlp-go/aibot"
)

func ast(cl *aibot.Client) {
	var s scanner.Scanner
	s.Init(strings.NewReader(`
	// comment
	id = "abcd" | id = 1234
	`))
	var err error
	s.Error = func(s *scanner.Scanner, msg string) {
		err = fmt.Errorf("%s (%s)", msg, s.Position)
	}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if err != nil {
			exitWithError(err)
		}
		fmt.Printf("%s '%s' '%s'\n", s.Position, scanner.TokenString(tok), s.TokenText())
	}
}

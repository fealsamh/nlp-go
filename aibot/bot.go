package aibot

import (
	"errors"
	"fmt"
)

type Bot struct {
	ID        string            `yaml:"id" json:"id,omitempty"`
	Name      string            `yaml:"name" json:"name,omitempty"`
	InitState string            `yaml:"initState" json:"initState,omitempty"`
	States    map[string]*State `taml:"states" json:"states,omitempty"`
}

type State struct {
	Message *Message `yaml:"message" json:"message,omitempty"`
	Action  string   `yaml:"action" json:"action,omitempty"`
}

type Message struct {
	Text string `yaml:"text" json:"text,omitempty"`
}

func (b *Bot) Validate() error {
	if b.States == nil {
		return errors.New("no states defined")
	}

	if b.InitState == "" {
		b.InitState = "start"
	}
	if _, ok := b.States[b.InitState]; !ok {
		return fmt.Errorf("initial state '%s' not found", b.InitState)
	}

	for sid, s := range b.States {
		if s.Message == nil {
			return fmt.Errorf("missing message in state '%s'", sid)
		}
		if s.Action == "" {
			return fmt.Errorf("missing action in state '%s'", sid)
		}
	}

	return nil
}

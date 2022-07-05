package aibot

import "fmt"

type Bot struct {
	ID        string            `yaml:"id" json:"id,omitempty"`
	Name      string            `yaml:"name" json:"name,omitempty"`
	InitState string            `yaml:"initState" json:"initState,omitempty"`
	States    map[string]*State `taml:"states" json:"states,omitempty"`
}

type State struct {
	Description string `yaml:"description" json:"description,omitempty"`
	Action      string `yaml:"action" json:"action,omitempty"`
}

func (b *Bot) Validate() error {
	if _, ok := b.States[b.InitState]; !ok {
		return fmt.Errorf("initial state '%s' not found", b.InitState)
	}

	return nil
}

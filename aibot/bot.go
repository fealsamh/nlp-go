package aibot

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Bot struct {
	ID        string             `yaml:"id,omitempty" json:"id,omitempty"`
	Name      string             `yaml:"name,omitempty" json:"name,omitempty"`
	Language  string             `yaml:"language,omitempty" json:"language,omitempty"`
	InitState string             `yaml:"init_state,omitempty" json:"init_state,omitempty"`
	Intents   map[string]*Intent `yaml:"intents,omitempty" json:"intents,omitempty"`
	States    map[string]*State  `yaml:"states,omitempty" json:"states,omitempty"`
	Entities  map[string]*Entity `yaml:"entities,omitempty" json:"entities,omitempty"`
	Synonyms  []string           `yaml:"synonyms,omitempty" json:"synonyms,omitempty"`
}

type Entity struct {
	Values []string `yaml:"values,omitempty" json:"values,omitempty"`
}

type Intent struct {
	Examples []string `yaml:"examples,omitempty" json:"examples,omitempty"`
}

type State struct {
	Message *Message `yaml:"message,omitempty" json:"message,omitempty"`
	Action  string   `yaml:"action,omitempty" json:"action,omitempty"`
	Next    []*Next  `yaml:"next,omitempty" json:"next,omitempty"`
}

type Message struct {
	Text    string   `yaml:"text,omitempty" json:"text,omitempty"`
	Options []string `yaml:"options,omitempty" json:"options,omitempty"`
}

type Next struct {
	State  string `yaml:"state,omitempty" json:"state,omitempty"`
	Option string `yaml:"option,omitempty" json:"option,omitempty"`
}

func (m *Message) HasOption(o string) bool {
	for _, o2 := range m.Options {
		if o == o2 {
			return true
		}
	}
	return false
}

func (b *Bot) Validate() error {
	if b.ID == "" {
		return errors.New("no bot ID provided")
	}

	if b.Language == "" {
		return errors.New("no bot language provided")
	}

	if b.Intents == nil {
		return errors.New("no intents defined")
	}

	for k, e := range b.Entities {
		if e == nil || len(e.Values) == 0 {
			return fmt.Errorf("no values for entity '%s'", k)
		}
	}

	for iid, i := range b.Intents {
		if err := i.Validate(iid); err != nil {
			return err
		}
	}

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
		if err := s.Validate(sid, b); err != nil {
			return err
		}
	}

	return nil
}

func (i *Intent) Validate(iid string) error {
	if len(i.Examples) == 0 {
		return fmt.Errorf("no examples for intent '%s'", iid)
	}
	return nil
}

func (s *State) Validate(sid string, b *Bot) error {
	if s.Message == nil {
		return fmt.Errorf("missing message in state '%s'", sid)
	}
	if s.Action == "" {
		return fmt.Errorf("missing action in state '%s'", sid)
	}
	switch s.Action {
	case "intent":
		if len(s.Next) > 0 {
			return fmt.Errorf("no next states expected in intent state '%s'", sid)
		}
	case "goto":
		if len(s.Next) == 0 {
			return fmt.Errorf("missing next states in goto state '%s'", sid)
		}
	default:
		return fmt.Errorf("unknown action '%s' in state '%s'", s.Action, sid)
	}
	for _, n := range s.Next {
		if n.Option != "" && !s.Message.HasOption(n.Option) {
			return fmt.Errorf("option '%s' not defined in state '%s'", n.Option, sid)
		}
		if _, ok := b.States[n.State]; !ok {
			return fmt.Errorf("next state '%s' not defined (used in state '%s')", n.State, sid)
		}
	}
	return nil
}

func (bot *Bot) Value() (driver.Value, error) {
	return json.Marshal(bot)
}

func (bot *Bot) Scan(data interface{}) error {
	b, ok := data.([]byte)
	if !ok {
		return errors.New("failed to scan, expected []byte")
	}
	return json.Unmarshal(b, &bot)
}

package aibot

type Bot struct {
	ID        string            `yaml:"id" json:"id"`
	Name      string            `yaml:"name" json:"name"`
	InitState string            `yaml:"initState" json:"initState"`
	States    map[string]*State `taml:"states" json:"states"`
}

type State struct {
	Description string `yaml:"description" json:"description"`
}

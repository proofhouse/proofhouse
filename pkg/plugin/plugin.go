package plugin

import "github.com/pkg/errors"

type Handle func(p Params)

type Step struct {
	text     string
	key      string
	argNames []string
	handle   Handle
}

func (s *Step) Text() string       { return s.text }
func (s *Step) Key() string        { return s.key }
func (s *Step) ArgNames() []string { return s.argNames }
func (s *Step) Handle() Handle     { return s.handle }

type Plugin struct {
	name  string
	steps map[string]Handle
}

// New creates plugin struct with given name.
func New(name string) *Plugin {
	return &Plugin{
		name:  name,
		steps: make(map[string]Handle),
	}
}

// AddStep adds new step to the plugin.
// text is step definition, for instance:
//
//     I send POST request to :url
//
// f is a pointer to function which will be called by Proofhouse:
//
//     func (url string) {}
func (p *Plugin) AddStep(text string, f Handle) {
	if f == nil {
		panic(errors.Errorf("Failed to add step to plugin '%v': step's function cannot be nil", p.name))
	}
	if _, dup := p.steps[text]; dup {
		panic(errors.Errorf("Failed to add step to plugin '%v': step '%v' was added previously", p.name, text))
	}

	p.steps[text] = f
}

// Name returns plugin unique name.
func (p *Plugin) Name() string {
	return p.name
}

// Steps returns all registered plugin steps.
func (p *Plugin) Steps() map[string]Handle {
	return p.steps
}

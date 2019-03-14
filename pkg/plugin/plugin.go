package plugin

import "github.com/pkg/errors"

type stepFunc func(p Params)

type Plugin struct {
	name  string
	steps map[string]stepFunc
}

// New creates plugin struct with given name.
func New(name string) *Plugin {
	return &Plugin{
		name:  name,
		steps: make(map[string]stepFunc),
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
func (p *Plugin) AddStep(text string, f stepFunc) {
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
func (p *Plugin) Steps() map[string]stepFunc {
	return p.steps
}

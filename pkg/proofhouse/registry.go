package proofhouse

import (
	"github.com/pkg/errors"
	"strings"
	"sync"
)

type Initializer func()

// Plugin registry struct.
type Registry struct {
	sync.RWMutex
	plugins      map[string]*Plugin
	initializers map[string]Initializer
	steps        map[string]Step
}

var registry = Registry{
	plugins:      make(map[string]*Plugin),
	initializers: make(map[string]Initializer),
	steps:        make(map[string]Step),
}

// Register registers provided pluginpointers in the registry under unique name.
func Register(name string, initializer Initializer) {
	registry.add(name, initializer)
}

// GetRegistry returns registry struct filled with plugins.
func GetRegistry() *Registry {
	return &registry
}

// List returns plugins map.
func (r *Registry) List() map[string]*Plugin {
	return r.plugins
}

// Step looks for step definition for given step key.
func (r *Registry) Step(key string) (step Step, err error) {
	step, ok := r.steps[key]
	if !ok {
		err = errors.Errorf("No step found for key: %v", key)
	}

	return
}

// add provided pluginpointers under given name.
func (r *Registry) add(name string, initializer Initializer) {
	r.Lock()
	defer r.Unlock()

	r.initializers[name] = initializer
}

type ParsedStepText struct {
	key  string
	args []string
}

func parseStepText(text string) ParsedStepText {
	var str strings.Builder
	var argsBuf strings.Builder
	var args []string

	skip := false
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		var nextRune rune
		if i+1 < len(runes) {
			nextRune = runes[i+1]
		}

		if skip {
			if runes[i] != ' ' {
				argsBuf.WriteRune(runes[i])
			}

			if runes[i] == ' ' || i+1 == len(runes) {
				str.WriteString("Ѻ")
				skip = false
				args = append(args, argsBuf.String())
				argsBuf.Reset()
				if runes[i] == ' ' {
					str.WriteRune(' ')
				}
			}

			continue
		} else if runes[i] == ':' && i+1 < len(runes) && nextRune != ' ' {
			skip = true
			continue
		}

		str.WriteRune(runes[i])
	}

	return ParsedStepText{
		key:  str.String(),
		args: args,
	}
}

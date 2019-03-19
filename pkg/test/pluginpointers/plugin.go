package pluginpointers

import (
	"fmt"
	"math/rand"
)

type Init

type PluginInitializer()

type StepArgs struct{}
type StepHandle func(args StepArgs)

type Plugin interface {
	GetName() string
	CreateHandler(int) Handler
}
type Handler interface {
	BeforeScenario()
	AfterScenario()
	StepHandles() map[string]StepHandle
}

type PluginImpl struct {
	name string
}

func NewPluginImpl(name string) Plugin {
	return &PluginImpl{
		name: name,
	}
}
func (p *PluginImpl) GetName() string { return p.name }
func (p *PluginImpl) CreateHandler(n int) Handler {
	return &HandlerImpl{
		plugin: p,
	}
}

type HandlerImpl struct {
	plugin Plugin
}

func (h *HandlerImpl) BeforeScenario() {}
func (h *HandlerImpl) AfterScenario()  {}
func (h *HandlerImpl) StepHandles() map[string]StepHandle {

	m := make(map[string]StepHandle)

	for i := 0; i < 20; i++ {
		m[h.plugin.na]
	}

	return map[string]StepHandle{
		fmt.Sprintf("I send POST requests to :url %v", rand.Float64()):              h.Step,
		fmt.Sprintf("I send GET requests to :url %v", rand.Float64()):               h.Step,
		fmt.Sprintf("I send PUT requests to :url %v", rand.Float64()):               h.Step,
		fmt.Sprintf("I send DELETE requests to :url %v", rand.Float64()):            h.Step,
		fmt.Sprintf("I send :num requests to :url %v", rand.Float64()):              h.Step,
		"I add HTTP header :name with value :value":                                 h.Step,
		fmt.Sprintf("I see JSON response: :response %v", rand.Float64()):            h.Step,
		fmt.Sprintf("I see HTTP header :name with value :value %v", rand.Float64()): h.Step,
		fmt.Sprintf("I dont see HTTP header :name %v", rand.Float64()):              h.Step,
		fmt.Sprintf("I see response is JSON %v", rand.Float64()):                    h.Step,
	}
}
func (h *HandlerImpl) Step(args StepArgs) {

}

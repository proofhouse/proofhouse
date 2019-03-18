package proofhouse

type Initializer func(config *Config) (plugin Plugin, steps map[string]Handle)
type Handle func(plugin *PluginBase)

type Plugin interface {
	BeforeScenario(ctx Context)
	AfterScenario(ctx Context)
}

type PluginBase struct{}

func (p *PluginBase) BeforeScenario(ctx Context) {}
func (p *PluginBase) AfterScenario(ctx Context)  {}

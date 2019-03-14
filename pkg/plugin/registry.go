package plugin

import (
	"sync"
)

// Plugin registry struct.
type Registry struct {
	sync.RWMutex
	plugins map[string]*Plugin
}

var registry = Registry{
	plugins: make(map[string]*Plugin),
}

// Register registers provided plugin in the registry under unique name.
func Register(plg *Plugin) {
	registry.add(plg.name, plg)
}

// GetRegistry returns registry struct filled with plugins.
func GetRegistry() *Registry {
	return &registry
}

// List returns plugins map.
func (r *Registry) List() map[string]*Plugin {
	return r.plugins
}

// add provided plugin under given name.
func (r *Registry) add(name string, plg *Plugin) {
	r.Lock()
	defer r.Unlock()

	if plg == nil {
		panic("Failed to register plugin: provided plugin object is nil")
	}
	if _, dup := r.plugins[name]; dup {
		panic("Failed to register plugin: plugin with name " + name + " already exists")
	}

	r.plugins[name] = plg
}

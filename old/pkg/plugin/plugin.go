// Package defines interfaces to be implemented by plugins as used by Proofhouse.
package plugin

// Plugin is the interface that must be implemented by Proofhouse plugins.
//
// Plugin struct is initialized ones for the whole application run cycle.
type Plugin interface {
	Name() string
	NewHandler() Handler
	Steps() map[string]interface{}
}

// Handler contains steps logic.
type Handler interface{}

type InitContext interface {
	Service() Service
}

type ExecutionContext interface{}

type StepArgs interface{}

type Service interface{}

// Base Plugin struct for future methods stub. Every plugin must extend this struct in case new interface method will
// be added.
type Base struct{}

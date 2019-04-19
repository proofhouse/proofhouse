package container

import (
	"github.com/proofhouse/proofhouse/pkg/service"
	"sync"
)

var (
	serviceFactoriesMu sync.Mutex
	serviceFactories = make(map[string]service.Factory)
)

func RegisterServiceFactory(factory service.Factory) {
	serviceFactoriesMu.Lock()
	defer serviceFactoriesMu.Unlock()

	serviceFactories[factory.Type()] = factory
}

func GetServiceFactory(t string) service.Factory {
	return serviceFactories[t]
}
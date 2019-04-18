package proofhouse

import "github.com/proofhouse/proofhouse/pkg/service"

var services = make(map[string]service.Service)

func RegisterService(service service.Service) {
	services[service.Type()] = service
}

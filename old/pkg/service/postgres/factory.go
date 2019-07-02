package postgres

import (
	"github.com/proofhouse/proofhouse/pkg/container"
	"github.com/proofhouse/proofhouse/pkg/service"
)

type Factory struct {

}

func (f *Factory) Type() string {
	return "postgres"
}

func (f *Factory) Kind() string {
	return "sql"
}

func (f *Factory) Create() service.Service {
	return &Service{}
}


func init() {
	container.RegisterServiceFactory(&Factory{})
}

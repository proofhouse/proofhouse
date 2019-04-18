package postgres

import "github.com/proofhouse/proofhouse/pkg/proofhouse"

type Postgres struct {
}

func (s *Postgres) Type() string {
	return "postgres"
}

func (s *Postgres) Kind() string {
	return "sql"
}

func init() {
	proofhouse.RegisterService(&Postgres{})
}

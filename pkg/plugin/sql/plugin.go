package sql

import (
	"database/sql"
	"github.com/cucumber/gherkin-go"
	"github.com/proofhouse/proofhouse/pkg/plugin"
	"github.com/proofhouse/proofhouse/pkg/proofhouse"
)

const pluginName = "sql"

type Plugin struct {
}

func (p *Plugin) Name() string {
	return pluginName
}

func (p *Plugin) Steps() map[string]interface{} {
	return map[string]interface{}{
		"Я вижу в базе данных запись :name: :table": (*Handler).SeeRecords,
	}
}

func (p *Plugin) NewHandler() plugin.Handler {
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=test sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Handler{
		db: db,
	}
}

type Handler struct {
	db *sql.DB
}

func (h *Handler) SeeRecords(name string, table gherkin.DataTable) {
	h.db.Exec("aga")
}

func init() {
	proofhouse.Register(&Plugin{})
}

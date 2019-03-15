package http

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/proofhouse/proofhouse/pkg/plugin"
)

func init() {
	connStr := "host=127.0.0.1 port=13100 user=postgres password=postgres dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_ = db

	p := plugin.New("http")
	p.AddStep("I send :num requests to :url", func(p plugin.Params) {
		url := p.String("url")
		fmt.Println("URL:", url, "NUM:", p.Int("num"))
	})

	plugin.Register(p)
}

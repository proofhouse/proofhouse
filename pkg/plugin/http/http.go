package http

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/proofhouse/proofhouse/pkg/proofhouse"
)

type Http struct {
	proofhouse.PluginBase
}

func (p *Http) Test() {
	fmt.Printf("AGAGAG TEST PLUGIN EEEE")
}

func da(handle proofhouse.Handle) {
	fmt.Printf("%T\n", handle)
}

func Test() {

}

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


	var handle proofhouse.Handle = (*Http).Test
	}

	da(handle)


	_ = db

	proofhouse.Register("http", func(config *proofhouse.Config) (plugin proofhouse.Plugin, steps map[string]proofhouse.Handle) {
		//steps = map[string]proofhouse.Handle {
		//	"aga": (*Plugin).Test,
		//}

		return &Http{}, nil
	})
}

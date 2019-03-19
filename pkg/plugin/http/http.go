package http

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/proofhouse/proofhouse/pkg/proofhouse"
)

type Http struct {
}

type Handler struct {
}

func (p *Http) GetHandler(num int) *Handler {
	return &Handler{}
}

func (h *Handler) Test(args proofhouse.Args, ctx proofhouse.Context) {

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

	http := Http{}
	handler := http.GetHandler(1)
	_ = handler

	proofhouse.Aga(Handler.Test)
}

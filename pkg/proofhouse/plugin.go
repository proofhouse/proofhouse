package proofhouse

import "fmt"

type Plugin interface {
	CreateHandler(int) *Handler
}

type Handler interface {
	BeforeScenario()
	AfterScenario()
}

type Handle func(p Plugin, args Args, ctx Context)

func Aga(handle Handle) {
	fmt.Printf("%+v\n", handle)
}

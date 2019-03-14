package http

import (
	"fmt"
	"github.com/proofhouse/proofhouse/pkg/plugin"
)

func init() {
	p := plugin.New("http")
	p.AddStep("I send :num requests to :url", func(p plugin.Params) {
		url := p.String("url")
		fmt.Println(url)
	})

	plugin.Register(p)
}

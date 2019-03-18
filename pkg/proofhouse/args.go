package proofhouse

import (
	"github.com/pkg/errors"
	"strconv"
)

type Args struct {
	data map[string]string
}

func NewArgs(data map[string]string) Args {
	return Args{
		data: data,
	}
}

// Int returns value of parameter with the given name in 'int' type. It panics if no such parameter exists or if
// conversion error occurs.
func (p *Args) Int(name string) int {
	s := p.String(name)
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(errors.Errorf("Step definition failed: parameter '%v' expected to be type of 'int', '%+v' given.", name, s))
	}

	return v
}

// String returns value of parameter with the given name. It panics if no param with such name found.
func (p *Args) String(name string) string {
	v, ok := p.data[name]
	if !ok {
		panic(errors.Errorf("Step definition failed: parameter '%v' not found in step's text definition", name))
	}

	return v
}

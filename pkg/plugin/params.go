package plugin

import (
	"github.com/pkg/errors"
	"strconv"
)

type Params struct {
	data map[string]string
}

// Int returns value of parameter with the given name in 'int' type. It panics if no such parameter exists or if
// conversion error occurs.
func (p *Params) Int(name string) int {
	s := p.String(name)
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(errors.Errorf("Step definition failed: parameter '%v' expected to be type of 'int', '%+v' given.", name, s))
	}

	return v
}

// String returns value of parameter with the given name. It panics if no param with such name found.
func (p *Params) String(name string) string {
	v, ok := p.data[name]
	if !ok {
		panic(errors.Errorf("Step definition failed: parameter '%v' not found in step's text definition", name))
	}

	return v
}

package proofhouse

import (
	"fmt"
	"sync"
)

// Runner receives Features and run them in parallel.
type Runner struct {
	config *Config
}

// NewRunner creates new runner struct.
func NewRunner(config *Config) *Runner {
	return &Runner{
		config: config,
	}
}

// Run receives Features from the channel and runs them in separate goroutines.
func (r *Runner) Run(ch <-chan Feature) {
	wg := sync.WaitGroup{}

	for {
		feature, more := <-ch
		if !more {
			break
		}

		for _, scenario := range feature.scenarios {
			wg.Add(1)
			go r.exec(scenario, &wg)
		}
	}

	wg.Wait()
}

// Exec executes single scenario.
func (r *Runner) exec(scenario *Scenario, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf(scenario.gherkin.Name)
}

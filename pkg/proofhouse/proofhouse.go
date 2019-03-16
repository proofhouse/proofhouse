package proofhouse

import "github.com/cucumber/gherkin-go"

// Feature represents whole .feature file.
type Feature struct {
	gherkin *gherkin.Feature
	scenarios []Scenario
}

// Background contains steps that run before each Scenario.
type Background struct {
	gherkin *gherkin.Background
	steps []Step
}

// Scenario contains steps for concrete test.
type Scenario struct {
	gherkin *gherkin.Scenario
	steps []Step
}

// Step contains necessary data to run a single step.
type Step struct {

}


package main

import (
	"bytes"
	"fmt"
	"github.com/cucumber/gherkin-go"
	"github.com/pkg/errors"
	"github.com/proofhouse/proofhouse/pkg/plugin"
	_ "github.com/proofhouse/proofhouse/pkg/plugin/http"
	"github.com/proofhouse/proofhouse/pkg/proofhouse"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Feature struct {
	path     string
	fileInfo os.FileInfo
}

func NewFeature(path string, fileInfo os.FileInfo) Feature {
	return Feature{
		path:     path,
		fileInfo: fileInfo,
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%+v", r)
			os.Exit(-1)
		}
	}()

	config, err := proofhouse.NewConfig("./configuration.yaml")
	if err != nil {
		panic(err)
	}

	var features []Feature

	err = filepath.Walk(config.FeaturesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "Failed to read feature file '%v'", path)
		}
		if info.IsDir() || filepath.Ext(path) != ".feature" {
			return nil
		}

		features = append(features, NewFeature(path, info))

		return nil
	})
	if err != nil {
		panic(err)
	}

	registry := plugin.GetRegistry()

	for _, f := range features {
		data, err := ioutil.ReadFile(f.path)
		if err != nil {
			panic(errors.Wrapf(err, "Failed to read feature file '%v'", f.path))
		}

		doc, err := gherkin.ParseGherkinDocument(bytes.NewReader(data))
		if err != nil {
			panic(errors.Wrap(err, "Failed to parse feature"))
		}

		for _, s := range doc.Feature.Children {
			switch s := s.(type) {
			default:
				panic(errors.Errorf("Unexpected type '%T'", s))
			case *gherkin.Scenario:
				runScenario(s, registry)
			case *gherkin.ScenarioOutline:
				fmt.Println("OUTLINE!!!")
			}
		}
	}

}

func runScenario(scenario *gherkin.Scenario, registry *plugin.Registry) {
	for _, step := range scenario.Steps {
		fmt.Println("Original text: ", step.Text)
		fmt.Println("Unified text: ", unifyStepText(step.Text))
	}
}

func unifyStepText(text string) string {
	var str strings.Builder

	for _, r := range text {
		str.WriteRune(r)
	}

	return str.String()
}

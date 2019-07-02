package proofhouse

import (
	"bytes"
	"github.com/cucumber/gherkin-go"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Parser scans features directory and all subdirectories for Feature files and creates ready to run struct
// for the runner.
type Parser struct {
	config *Config
}

// NewParser creates new Parser
func NewParser(config *Config) *Parser {
	return &Parser{
		config: config,
	}
}

// Parse locates feature files, parses them to Feature struct and sends them to the channel "ch".
func (p *Parser) Parse(ch chan<- Feature) {
	err := filepath.Walk(p.config.FeaturesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "Failed to read feature file '%v'", path)
		}
		if info.IsDir() || filepath.Ext(path) != ".feature" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		doc, err := gherkin.ParseGherkinDocument(bytes.NewReader(data))
		if err != nil {
			return err
		}

		ch <- Feature{
			gherkin:  doc.Feature,
			fileInfo: info,
		}

		return nil
	})

	close(ch)

	if err != nil {
		panic(err)
	}
}

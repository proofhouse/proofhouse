package proofhouse

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


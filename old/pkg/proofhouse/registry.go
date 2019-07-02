package proofhouse

import (
	"github.com/proofhouse/proofhouse/pkg/plugin"
	"strings"
	"sync"
)

var (
	pluginsMu sync.RWMutex
	plugins   = make(map[string]plugin.Plugin)
)

func Register(plugin plugin.Plugin) {
	pluginsMu.Lock()
	defer pluginsMu.Unlock()

	plugins[plugin.Name()] = plugin
}

type ParsedStepText struct {
	key  string
	args []string
}

func parseStepText(text string) ParsedStepText {
	var str strings.Builder
	var argsBuf strings.Builder
	var args []string

	skip := false
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		var nextRune rune
		if i+1 < len(runes) {
			nextRune = runes[i+1]
		}

		if skip {
			if runes[i] != ' ' {
				argsBuf.WriteRune(runes[i])
			}

			if runes[i] == ' ' || i+1 == len(runes) {
				str.WriteString("Ѻ")
				skip = false
				args = append(args, argsBuf.String())
				argsBuf.Reset()
				if runes[i] == ' ' {
					str.WriteRune(' ')
				}
			}

			continue
		} else if runes[i] == ':' && i+1 < len(runes) && nextRune != ' ' {
			skip = true
			continue
		}

		str.WriteRune(runes[i])
	}

	return ParsedStepText{
		key:  str.String(),
		args: args,
	}
}

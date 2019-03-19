package test

import (
	"github.com/proofhouse/proofhouse/pkg/test/pluginpointers"
	"strconv"
	"testing"
)

func BenchmarkPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		app := pluginpointers.NewApp()
		for i := 0; i < 20; i++ {
			app.RegisterPlugin(pluginpointers.NewPluginImpl("plugin_" + strconv.Itoa(i)))
		}

		app.Run()

		app = nil
	}
}

package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestStepRecognitionReflection(t *testing.T) {
	steps := make(map[string]interface{})
	steps["I send :num requests to :url"] = func(num int, url string) {

	}

	r := reflect.ValueOf(steps["I send :num requests to :url"])
	t.Fatal(r.Type().NumIn())

	text := `I send 12 requests to "https://example.org"`
	_ = text
}

func unifyStepText1(text string) string {
	var str strings.Builder

	skip := false
	for _, r := range text {
		if r == '"' {
			if skip {
				skip = false
				str.WriteString("_var_")
				continue
			} else {
				skip = true
			}
		}

		if skip {
			if r == '"' {

			}
		}

		str.WriteRune(r)
	}

	return str.String()
}

func unifyStepText2(text string) string {
	r, _ := regexp.Compile(`".*?"`)

	return r.ReplaceAllString(text, "XAXA")
}

func TestAga(t *testing.T) {
	t.Fatal(unifyStepText1(`I send "2" requests to "https://example.org"`))
}

func BenchmarkNewFeatureAga1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unifyStepText1(`I send "2" requests to "https://example.org"`)
	}
}

func BenchmarkNewFeatureAga2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unifyStepText2(`I send "2" requests to "https://example.org"`)
	}
}

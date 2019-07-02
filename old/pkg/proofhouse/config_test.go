package proofhouse

import "testing"

func TestConfigReturnNilOnError(t *testing.T) {
	c, _ := NewConfig("testdata/somenonexistent.yaml")
	if c != nil {
		t.Errorf("c != nil")
	}
}

func TestConfigNotExists(t *testing.T) {
	_, err := NewConfig("testdata/somenonexistent.yaml")
	if err == nil {
		t.Error("err == nil")
	}
	if err.Error() != "Failed to read configuration file 'testdata/somenonexistent.yaml': open testdata/somenonexistent.yaml: no such file or directory" {
		t.Errorf("err.Error() != 'Failed to read configuration file 'testdata/somenonexistent.yaml': open somenonexistent.yaml: no such file or directory', '%v' given", err.Error())
	}
}

func TestConfigInvalidFormat(t *testing.T) {
	_, err := NewConfig("../../testdata/invalid_configuration.yaml")
	if err.Error() != "Error occurred while parsing yaml configuration file: yaml: unmarshal errors:\nline 1: cannot unmarshal !!str `invalid...` into proofhouse.Config" {
		t.Errorf("err.Error() != '', '%v' given", err.Error())
	}
}
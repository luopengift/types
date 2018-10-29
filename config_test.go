package types

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type TestConfig struct {
	XMLName xml.Name `json:"-" yaml:"-" xml:"testconfig"`
	Name    string   `json:"name" yaml:"name" xml:"name"`
	Number  int      `json:"number" yaml:"number" xml:"number"`
	List    []string `json:"list" yaml:"list" xml:"list"`
}

type TestConfigINI struct {
	Testconfig struct {
		Name   string
		Number int
		List   []string
	}
}

func Test_JSON(t *testing.T) {
	config := TestConfig{}
	err := ParseConfigFile(&config, "test/config.json")
	fmt.Println("JSON:", err, config)
}

func Test_YAML(t *testing.T) {
	config := TestConfig{}
	err := ParseConfigFile(&config, "test/config.yaml")
	fmt.Println("YAML:", err, config)
}

func Test_XML(t *testing.T) {
	config := TestConfig{}
	err := ParseConfigFile(&config, "test/config.xml")
	fmt.Println("XML:", err, config)
}

func Test_INI(t *testing.T) {
	config := TestConfigINI{}
	err := ParseConfigFile(&config, "test/config.ini")
	fmt.Println("INI:", err, config)
}

// Configer config interface.
type Configer interface {
	Parse(v interface{}) error
}

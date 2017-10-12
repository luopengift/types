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
	err := ParseFile("test/config.json", &config)
	fmt.Println("JSON:", err, config)
}

func Test_YAML(t *testing.T) {
	config := TestConfig{}
	err := ParseFile("test/config.yaml", &config)
	fmt.Println("YAML:", err, config)
}

func Test_XML(t *testing.T) {
	config := TestConfig{}
	err := ParseFile("test/config.xml", &config)
	fmt.Println("XML:", err, config)
}

func Test_INI(t *testing.T) {
	config := TestConfigINI{}
	err := ParseFile("test/config.ini", &config)
	fmt.Println("INI:", err, config)
}

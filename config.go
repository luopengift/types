package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/gcfg.v1"
	"gopkg.in/yaml.v2"
)

// FormatJSON format json
func FormatJSON(in, out interface{}) error {
	var err error
	if b, err := ToBytes(in); err == nil {
		return json.Unmarshal(b, out)
	}
	return err
}

// FormatXML format xml
func FormatXML(in, out interface{}) error {
	var err error
	if b, err := ToBytes(in); err == nil {
		return xml.Unmarshal(b, out)
	}
	return err
}

// FormatYAML format yaml
func FormatYAML(in, out interface{}) error {
	var err error
	if b, err := ToBytes(in); err == nil {
		return yaml.Unmarshal(b, out)
	}
	return err
}

// FormatINI format ini
func FormatINI(in, out interface{}) error {
	var err error
	if s, err := ToString(in); err == nil {
		return gcfg.ReadStringInto(out, s)
	}
	return err
}

// ParseConfigFile parse config file
func ParseConfigFile(v interface{}, file string) error {
	filepath := strings.Replace(file, "~", os.Getenv("HOME"), -1)
	b, err := FileToBytes(filepath)
	if err != nil {
		return err
	}
	switch suffix := path.Ext(filepath); suffix {
	case ".json", ".js":
		return FormatJSON(b, v)
	case ".xml":
		return FormatXML(b, v)
	case ".ini":
		return FormatINI(b, v)
	case ".yaml", ".yml":
		return FormatYAML(b, v)
	case ".conf":
		fallthrough
	default:
		return fmt.Errorf("unknown suffix: %s", suffix)
	}
}

// ParseConfigFiles parse config files
func ParseConfigFiles(v interface{}, files ...string) error {
	for _, file := range files {
		if err := ParseConfigFile(v, file); err != nil {
			return err
		}
	}
	return nil
}

//ToYAML 将结构体in转换成yml格式的字符串, 适用于配置文件落地
func ToYAML(in interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

// ToJSON to json
func ToJSON(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

// ToXML to json
func ToXML(in interface{}) ([]byte, error) {
	return xml.Marshal(in)
}

// Configer config interface.
type Configer interface {
	Parse(v interface{}) error
}

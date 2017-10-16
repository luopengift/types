package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/gcfg.v1"
	"gopkg.in/yaml.v2"
	"path"
)

func FormatJSON(in, out interface{}) error {
	var err error
	if b, err := ToBytes(in); err == nil {
		return json.Unmarshal(Clean(b), out)
	}
	return err
}

func FormatXML(in, out interface{}) error {
	var err error
	if b, err := ToBytes(in); err == nil {
		return xml.Unmarshal(b, out)
	}
	return err
}

func FormatYAML(in, out interface{}) error {
	var err error
	if b, err := ToBytes(in); err == nil {
		return yaml.Unmarshal(b, out)
	}
	return err
}

func FormatINI(in, out interface{}) error {
	var err error
	if s, err := ToString(in); err == nil {
		return gcfg.ReadStringInto(out, s)
	}
	return err
}

func ParseConfigFile(file string, v interface{}) error {
	b, err := FileToBytes(file)
	if err != nil {
		return err
	}
	switch suffix := path.Ext(file); suffix {
	case ".json":
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
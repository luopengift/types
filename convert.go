package types

import (
    "encoding/json"
)

// in format out
// eg: Format(map[string]interface{...}, &Struct{})
func Format(in, out interface{}) error {
    var err error
    if b, err := json.Marshal(in); err == nil {
        err = json.Unmarshal(b, out)
    }
    return err
}

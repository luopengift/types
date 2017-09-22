package types

import (
	"github.com/luopengift/golibs/logger"
	"testing"
)

func Test_Int(t *testing.T) {
	logger.Warn("==============")
	var i Int = 1
	logger.Info("defined type: %#v<%#T>", i, i)
	logger.Info("convert int: %#v<%#T>", i.Int(), i.Int())
	logger.Info("convert int64: %#v<%#T>", i.Int64(), i.Int64())
	logger.Info("convert float64: %#v<%#T>", i.Float64(), i.Float64())
	logger.Info("convert string: %#v<%#T>", i.String(), i.String())
	logger.Info("convert []byte: %#v<%#T>", i.Bytes(), i.Bytes())
}
func Test_Int64(t *testing.T) {
	logger.Warn("==============")
	var i Int64 = 1
	logger.Info("defined type: %#v<%#T>", i, i)
	logger.Info("convert int: %#v<%#T>", i.Int(), i.Int())
	logger.Info("convert int64: %#v<%#T>", i.Int64(), i.Int64())
	logger.Info("convert float64: %#v<%#T>", i.Float64(), i.Float64())
	logger.Info("convert string: %#v<%#T>", i.String(), i.String())
	logger.Info("convert []byte: %#v<%#T>", i.Bytes(), i.Bytes())
}
func Test_String(t *testing.T) {
	logger.Warn("==============")
	var i String = "1.1"
	logger.Info("defined type: %#v<%#T>", i, i)
	logger.Info("convert int: %#v<%#T>", i.Int(), i.Int())
	logger.Info("convert int64: %#v<%#T>", i.Int64(), i.Int64())
	logger.Info("convert float64: %#v<%#T>", i.Float64(), i.Float64())
	logger.Info("convert string: %#v<%#T>", i.String(), i.String())
	logger.Info("convert []byte: %#v<%#T>", i.Bytes(), i.Bytes())
}
func Test_Bytes(t *testing.T) {
	logger.Warn("==============")
	var i Bytes = []byte("1")
	logger.Info("defined type: %#v<%#T>", i, i)
	logger.Info("convert int: %#v<%#T>", i.Int(), i.Int())
	logger.Info("convert int64: %#v<%#T>", i.Int64(), i.Int64())
	logger.Info("convert string: %#v<%#T>", i.String(), i.String())
	logger.Info("convert []byte: %#v<%#T>", i.Bytes(), i.Bytes())
}

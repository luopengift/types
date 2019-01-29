package types

import "time"

// Typer type interface
type Typer interface {
	Int() int
	Int64() int64
	Float64() float64
	String() string
}

var _ Typer = Int(1)
var _ Typer = Int64(1)
var _ Typer = Float64(1.0)
var _ Typer = String("1")
var _ Typer = Time(time.Now())

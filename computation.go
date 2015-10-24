package concord

import (
	_ "github.com/concord/concord-go/thrift"
)

// Computation defines Concord Computation.
type Computation interface {
	Init(*Context) error
	ProcessRecords(*Context, interface{}) error
	ProcessTimer(*Context, int64, string) error
	Metadata()
}

// Metadata holds metadata for Computation service.
type Metadata struct {
	Name            string
	Inputs, Outputs []Stream
}

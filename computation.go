package concord

import (
	_ "github.com/concord/concord-go/thrift"
)

// Computation defines Concord Computation.
type Computation interface {
	Init(*Context) error
	ProcessRecords(*Context, *Record) error
	ProcessTimer(*Context, int64, string) error
	Metadata() *Metadata
}

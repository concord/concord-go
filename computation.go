// Package concord implements API client library for
// Concord streams event processing system.
//
// Your worker must implement Computation interface.
// In main() handler call concord.Serve() to start the worker.
//
// Example:
//    package main
//
//    import "github.com/concord/concord-go"
//
//    type Demo struct{}
//    func (*Demo) Init(*concord.Context) error {return nil}
//    func (*Demo) ProcessRecords(*concord.Context, *concord.Record) error {return nil}
//    func (*Demo) ProcessTimer(*concord.Context, int64, string) error {return nil}
//    func (*Demo) Metadata() *concord.Metadata {
//        return &concord.Metadata{
//            Name: "demo",
//        }
//    }
//
//    func main() {
//        concord.Serve(&Demo{})
//    }
//
// See more real-world examples in examples/ subdirectory.
package concord

// Computation defines Concord Computation.
//
// User must implement this interface for it's own computation worker.
type Computation interface {
	Init(*Context) error
	ProcessRecords(*Context, *Record) error
	ProcessTimer(*Context, int64, string) error
	Metadata() *Metadata
}

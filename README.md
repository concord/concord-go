# concord-go
Concord Go(lang) client

[![GoDoc](https://godoc.org/github.com/concord/concord-go?status.svg)](https://godoc.org/github.com/concord/concord-go)

### Installation
```
go get -u github.com/concord/concord-go
```

### Usage
Make sure you're familiar with Concord concepts. The best place to start is to read [Concord docs](http://docs.concord.io/concepts.html).


Then, implement interface `concord.Computation` to be used within Concord:

```go
type Computation interface {
    // Perform initialization of computation.
    Init(*Context) error

	// Perform some cleanup.
    Destroy() error

    // Process record from stream.
    ProcessRecords(*Context, *Record) error

    // Typically used to periodically write record to stream.
    ProcessTimer(*Context, int64, string) error

    // Defines name, input and output streams of computation.
    Metadata() *Metadata
}
```
  
In func `main` create instance `comp` of your Computation and invoke ```concord.Serve(comp)``` to start.

### Documentation

See [GoDoc documentation](https://godoc.org/github.com/concord/concord-go)

### Examples

See examples in [examples/](https://github.com/concord/concord-go/tree/master/examples) directory.

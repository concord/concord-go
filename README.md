# concord-go
Concord Go(lang) client

### Installation
```
go get github.com/concord/concord-go
```

### Usage
Implement interface ```concord.Computation``` in order to run computation on Concord.

```
type Computation interface {

  // Perform initialization of computation.
  Init(*Context) error

  // Process record from stream.
  ProcessRecords(*Context, *Record) error

  // Typically used to periodically write record to stream.
  ProcessTimer(*Context, int64, string) error

  // Defines name, input and output streams of computation.
  Metadata() *Metadata

}
```


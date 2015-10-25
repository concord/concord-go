# concord-go
Concord Go(lang) client

### Installation
```
go get github.com/concord/concord-go
```

### Usage
To run computation on Concord:
  1. Implement interface ```concord.Computation```
  ```go
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
  1. In func main create instance ```comp```that implements ```concord.Computation```
  1. Invoke ```concord.Serve(comp)```




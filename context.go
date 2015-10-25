package concord

import (
	bolt "github.com/concord/concord-go/thrift"
	"time"
)

// Context represents single Computation Context
type Context struct {
	tx *bolt.ComputationTx
}

// NewContext returns new Context.
func NewContext() *Context {
	tx := bolt.NewComputationTx()
	tx.Timers = make(map[string]int64)
	return &Context{
		tx: tx,
	}
}

// SetTimer sets 'name' timer to time 't'.
func (c *Context) SetTimer(t time.Time, name string) {
	t1 := t.UnixNano() / int64(time.Millisecond)
	c.tx.Timers[name] = t1
}

// ProduceRecord stores a record to be sent downstream to context.
func (c *Context) ProduceRecord(stream, key, value string) {
	record := bolt.NewRecord()
	record.Time = time.Now().UnixNano() / int64(time.Millisecond)
	record.Key = []byte(key)
	record.Data = []byte(value)
	record.UserStream = []byte(stream)
	record.Meta = bolt.NewRecordMetadata()

	c.tx.Records = append(c.tx.Records, record)
}

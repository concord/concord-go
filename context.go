package concord

import (
	bolt "github.com/concord/concord-go/thrift"
	"time"
)

type Context struct {
	Tx *bolt.ComputationTx
}

func NewContext() *Context {
	return &Context{
		Tx: bolt.NewComputationTx(),
	}
}

func (c *Context) SetTimer(t time.Time, name string) {
	t1 := t.UnixNano() / 1000 // milliseconds
	c.Tx.Timers[name] = t1
}

// ProduceRecord stores a record to be sent downstream to context.
func (c *Context) ProduceRecord(stream, key, value string) {
	record := bolt.NewRecord()
	record.Time = time.Now().UnixNano() / 1000
	record.Key = []byte(key)
	record.Data = []byte(value)
	record.UserStream = []byte(stream)

	c.Tx.Records = append(c.Tx.Records, record)
}

func (*Context) SetState(key string, value interface{}) {
	// TODO

}

func (*Context) GetState(key string) (interface{}, bool) {
	// TODO
	return nil, false
}

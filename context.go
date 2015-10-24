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

func (*Context) SetTimer(t time.Time, name string) {
	t1 := t.UnixNano() / 1000 // milliseconds
	// TODO
	_ = t1
}

func (*Context) ProduceRecord(name, key, value string) error {
	// TODO
	return nil
}

func (*Context) SetState(key string, value interface{}) {
	// TODO

}

func (*Context) GetState(key string) (interface{}, bool) {
	// TODO
	return nil, false
}

package concord

import (
	"errors"
	bolt "github.com/concord/concord-go/thrift"
	"log"
)

// ComputationService implements thrift's ComputationService && MutableEphemeralStateService
type ComputationService struct {
	comp  Computation
	proxy *Proxy
}

func NewComputationService(comp Computation, proxy *Proxy) *ComputationService {
	return &ComputationService{
		comp:  comp,
		proxy: proxy,
	}
}

func (*ComputationService) Init() (r *bolt.ComputationTx, err error) {
	ctx := NewContext()
	return ctx.Tx, nil
}

func (c *ComputationService) BoltProcessRecords(records []*bolt.Record) (r []*bolt.ComputationTx, err error) {
	var txs []*bolt.ComputationTx
	for _, record := range records {
		ctx := NewContext()
		err := c.comp.ProcessRecords(ctx, record)
		if err != nil {
			log.Println("[ERROR] error processing record:", err)
			return nil, errors.New("failed to process records")
		}
		txs = append(txs, ctx.Tx)
	}
	return txs, nil
}

func (c *ComputationService) BoltProcessTimer(key string, time int64) (r *bolt.ComputationTx, err error) {
	ctx := NewContext()

	if err := c.comp.ProcessTimer(ctx, time, key); err != nil {
		log.Println("[ERROR] error processing timer:", err)
		return nil, errors.New("failed to process timer")
	}
	return ctx.Tx, nil
}

func (c *ComputationService) BoltMetadata() (r *bolt.ComputationMetadata, err error) {
	return nil, nil
}

func (c *ComputationService) GetState(key string) ([]byte, error) {
	return c.proxy.GetState(key)
}

func (c *ComputationService) SetState(key string, value []byte) error {
	return c.proxy.SetState(key, value)
}

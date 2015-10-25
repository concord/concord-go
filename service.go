package concord

import (
	"errors"
	bolt "github.com/concord/concord-go/thrift"
	"log"
)

var (
	ErrProcessRecords = errors.New("failed to process records")
	ErrProcessTimer   = errors.New("failed to process timer")
)

// ComputationService is a high-level wrapper around Computation. It implements thrift's ComputationService && MutableEphemeralStateService
type computationService struct {
	comp  Computation
	proxy *proxy
}

// newComputationService returns new CompurationService wrapper object.
func newComputationService(comp Computation, proxy *proxy) *computationService {
	return &computationService{
		comp:  comp,
		proxy: proxy,
	}
}

// Init implements ComputationService.
func (c *computationService) Init() (*bolt.ComputationTx, error) {
	ctx := NewContext()
	err := c.comp.Init(ctx)
	return ctx.tx, err
}

// BoltProcessRecords implements ComputationService.
func (c *computationService) BoltProcessRecords(records []*bolt.Record) ([]*bolt.ComputationTx, error) {
	var txs []*bolt.ComputationTx
	for _, record := range records {
		ctx := NewContext()
		err := c.comp.ProcessRecords(ctx, &Record{*record})
		if err != nil {
			log.Println("[ERROR] error processing record:", err)
			return nil, ErrProcessRecords
		}
		txs = append(txs, ctx.tx)
	}
	return txs, nil
}

// BoltProcessTimer implements ComputationService.
func (c *computationService) BoltProcessTimer(key string, time int64) (*bolt.ComputationTx, error) {
	ctx := NewContext()

	if err := c.comp.ProcessTimer(ctx, time, key); err != nil {
		log.Println("[ERROR] error processing timer:", err)
		return nil, ErrProcessTimer
	}
	return ctx.tx, nil
}

// BoltMetadata implements ComputationService.
func (c *computationService) BoltMetadata() (*bolt.ComputationMetadata, error) {
	return nil, nil
}

// GetState implements MutableEphemeralStateService.
func (c *computationService) GetState(key string) ([]byte, error) {
	return c.proxy.GetState(key)
}

// SetState implements MutableEphemeralStateService.
func (c *computationService) SetState(key string, value []byte) error {
	return c.proxy.SetState(key, value)
}

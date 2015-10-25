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

// ComputationService implements thrift's ComputationService && MutableEphemeralStateService
type ComputationService struct {
	comp  Computation
	proxy *Proxy
}

// NewComputationService returns new CompurationService wrapper object.
func NewComputationService(comp Computation, proxy *Proxy) *ComputationService {
	return &ComputationService{
		comp:  comp,
		proxy: proxy,
	}
}

// Init implements ComputationService.
func (c *ComputationService) Init() (*bolt.ComputationTx, error) {
	ctx := NewContext()
	err := c.comp.Init(ctx)
	return ctx.tx, err
}

// BoltProcessRecords implements ComputationService.
func (c *ComputationService) BoltProcessRecords(records []*bolt.Record) ([]*bolt.ComputationTx, error) {
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
func (c *ComputationService) BoltProcessTimer(key string, time int64) (*bolt.ComputationTx, error) {
	ctx := NewContext()

	if err := c.comp.ProcessTimer(ctx, time, key); err != nil {
		log.Println("[ERROR] error processing timer:", err)
		return nil, ErrProcessTimer
	}
	return ctx.tx, nil
}

// BoltMetadata implements ComputationService.
func (c *ComputationService) BoltMetadata() (*bolt.ComputationMetadata, error) {
	return nil, nil
}

// GetState implements MutableEphemeralStateService.
func (c *ComputationService) GetState(key string) ([]byte, error) {
	return c.proxy.GetState(key)
}

// SetState implements MutableEphemeralStateService.
func (c *ComputationService) SetState(key string, value []byte) error {
	return c.proxy.SetState(key, value)
}

package concord

import (
	"errors"
	bolt "github.com/concord/concord-go/thrift"
	"log"
)

// ComputationService implements thrift's ComputationService
type ComputationService struct {
	comp Computation
}

func NewComputationService(comp Computation) *ComputationService {
	return &ComputationService{
		comp: comp,
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
		err := c.comp.ProcessRecords(ctx, &Record{*record})
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

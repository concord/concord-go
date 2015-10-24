package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

// ComputationService implements thrift's ComputationService
type ComputationService struct {
	comp *Computation
}

func NewComputationService(comp *Computation) *ComputationService {
	return &ComputationService{
		comp: comp,
	}
}

func (*ComputationService) Init() (r *bolt.ComputationTx, err error) {
	ctx := NewContext()
	return ctx.Tx, nil
}
func (*ComputationService) BoltProcessRecords(records []*bolt.Record) (r []*bolt.ComputationTx, err error) {
	return nil, nil
}
func (*ComputationService) BoltProcessTimer(key string, time int64) (r *bolt.ComputationTx, err error) {
	return nil, nil
}
func (*ComputationService) BoltMetadata() (r *bolt.ComputationMetadata, err error) {
	return nil, err
}

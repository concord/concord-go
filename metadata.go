package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

// Metadata holds metadata for Computation service.
type Metadata struct {
	Name    string
	Inputs  []Stream
	Outputs []string
}

func (m *Metadata) ToBoltMetadata() *bolt.ComputationMetadata {
	var inputs []*bolt.StreamMetadata
	for _, input := range m.Inputs {
		md := bolt.NewStreamMetadata()
		md.Name = input.Name
		md.Grouping = input.Grouping
		inputs = append(inputs, md)
	}
	return &bolt.ComputationMetadata{
		Name:     m.Name,
		Istreams: inputs,
		Ostreams: m.Outputs,
	}
}

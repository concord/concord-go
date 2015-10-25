package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

// Metadata holds metadata for Computation service.
// Inputs and Outputs defines streams Computation is
// working on.
type Metadata struct {
	Name    string
	Inputs  []Stream
	Outputs []string
}

// toBoltMetadata converts our Metadata into internal bolt's metadata.
func (m *Metadata) toBoltMetadata() *bolt.ComputationMetadata {
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

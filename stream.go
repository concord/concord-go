package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

// Grouping types. First one will be default.
const (
	Shuffle StreamGrouping = iota
	GroupBy
	RoundRobin
	Custom
)

type (
	// Stream implements stream info.
	Stream struct {
		Name     string
		Grouping StreamGrouping
	}

	// StreamGrouping is a proxy type for bolt.StreamGrouping
	StreamGrouping int
)

// NewDefaultStream creates stream with default grouping option.
func NewDefaultStream(name string) *Stream {
	return &Stream{
		Name: name,
	}
}

// NewStream creates stream with given grouping option.
func NewStream(name string, grouping StreamGrouping) *Stream {
	return &Stream{
		Name:     name,
		Grouping: grouping,
	}
}

var groupingMap = map[StreamGrouping]bolt.StreamGrouping{
	Shuffle:    bolt.StreamGrouping_SHUFFLE,
	GroupBy:    bolt.StreamGrouping_GROUP_BY,
	RoundRobin: bolt.StreamGrouping_ROUND_ROBIN,
	Custom:     bolt.StreamGrouping_CUSTOM,
}

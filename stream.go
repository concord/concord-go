package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

// Stream implements stream info.
type Stream struct {
	Name     string
	Grouping bolt.StreamGrouping
}

// NewDefaultStream creates stream with default grouping option.
func NewDefaultStream(name string) Stream {
	return Stream{
		Name:     name,
		Grouping: bolt.StreamGrouping_SHUFFLE,
	}
}

// NewGroupByStream creates stream with group by grouping option.
func NewGroupByStream(name string) Stream {
	return Stream{
		Name:     name,
		Grouping: bolt.StreamGrouping_GROUP_BY,
	}
}

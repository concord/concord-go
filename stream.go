package concord

import (
	"github.com/concord/concord-go/thrift"
)

// StreamGrouping constants.
const (
	StreamShuffle = bolt.StreamGrouping_SHUFFLE
	StreamGroupBy = bolt.StreamGrouping_GROUP_BY
)

// Stream implements stream info.
type Stream struct {
	Name     string
	Grouping bolt.StreamGrouping
}

// NewDefaultStream returns new stream with default grouping option.
func NewDefaultStream(name string) Stream {
	return Stream{
		Name:     name,
		Grouping: StreamShuffle,
	}
}

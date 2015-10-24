package concord

import (
	"github.com/concord/concord-go/thrift"
)

const (
	StreamShuffle = bolt.StreamGrouping_SHUFFLE
	StreamGroupBy = bolt.StreamGrouping_GROUP_BY
)

type Stream struct {
	Name     string
	Grouping bolt.StreamGrouping
}

func NewSuffleStream(name string) Stream {
	return Stream{
		Name:     name,
		Grouping: StreamShuffle,
	}
}

func NewGroupByStream(name string) Stream {
	return Stream{
		Name:     name,
		Grouping: StreamGroupBy,
	}
}

package concord

import (
	bolt "github.com/concord/concord-go/thrift"
	"reflect"
	"testing"
)

func TestMetadata(t *testing.T) {
	md := &Metadata{
		Name: "test",
		Inputs: []*Stream{
			NewGroupByStream("stream 1"),
			NewDefaultStream("stream 2"),
		},
		Outputs: []string{"output 1", "output 2"},
	}
	boltMd := md.toBoltMetadata()
	want := &bolt.ComputationMetadata{
		Name: "test",
		Istreams: []*bolt.StreamMetadata{
			&bolt.StreamMetadata{
				Name:     "stream 1",
				Grouping: bolt.StreamGrouping_GROUP_BY,
			},
			&bolt.StreamMetadata{
				Name:     "stream 2",
				Grouping: bolt.StreamGrouping_SHUFFLE,
			},
		},
		Ostreams: []string{"output 1", "output 2"},
	}
	if !reflect.DeepEqual(boltMd, want) {
		t.Fatalf("Metadata convert failed. Want %v, but got %v", want, boltMd)
	}
}

func BenchmarkMetadata(b *testing.B) {
	md := &Metadata{
		Name: "test",
		Inputs: []*Stream{
			NewGroupByStream("stream 1"),
			NewDefaultStream("stream 2"),
		},
		Outputs: []string{"output 1", "output 2"},
	}
	for i := 0; i < b.N; i++ {
		md.toBoltMetadata()
	}
}

package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

type Record struct {
	bolt.Record
}

func NewRecord() *Record {
	rec := bolt.NewRecord()
	rec.Meta = bolt.NewRecordMetadata()
	return &Record{*rec}
}

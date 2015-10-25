package concord

import (
	bolt "github.com/concord/concord-go/thrift"
)

// Record represents single record.
type Record struct {
	bolt.Record
}

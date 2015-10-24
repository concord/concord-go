package concord

// Computation defines Concord Computation.
type Computation interface {
	Init(*Context) error
	ProcessRecords(*Context, interface{}) error
	ProcessTimer(*Context, int64, string) error
	Metadata() *Metadata
}

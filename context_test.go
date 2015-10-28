package concord

import (
	"bytes"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := NewContext(&proxy{})
	if ctx == nil {
		t.Fatal("NewContext should return non-nil context")
	}
	if ctx.tx.Timers == nil {
		t.Fatal("NewContext should initialize timers")
	}

	t1 := time.Now()
	ctx.SetTimer(t1, "test 1")
	ctx.SetTimer(t1.Add(time.Second), "test 2")

	if len(ctx.tx.Timers) != 2 {
		t.Fatalf("Should be %d timers, but got %d", 2, len(ctx.tx.Timers))
	}

	if ctx.tx.Timers["test 1"] != t1.UnixNano()/int64(time.Millisecond) {
		t.Fatalf("Timer should set correct time in milliseconds since epoch. Want %d, but got %d", t1.UnixNano()/int64(time.Millisecond), ctx.tx.Timers["test 1"])
	}

	ctx.ProduceRecord("stream 1", "key", "value")
	ctx.ProduceRecord("stream 1", "key2", "value2")
	ctx.ProduceRecord("stream 3", "key1", "value1")

	if len(ctx.tx.Records) != 3 {
		t.Fatalf("Should %d records, but got %d", 3, len(ctx.tx.Records))
	}
	rec := ctx.tx.Records[2]
	if rec.Time == 0 {
		t.Fatal("Record time should be set")
	}
	if !bytes.Equal(rec.Key, []byte(`key1`)) {
		t.Fatalf("Record key should be %s, but got %s", "key1", string(rec.Key))
	}
	if !bytes.Equal(rec.Data, []byte(`value1`)) {
		t.Fatalf("Record data should be %s, but got %s", "value1", string(rec.Data))
	}
	if !bytes.Equal(rec.UserStream, []byte(`stream 3`)) {
		t.Fatalf("Record stream should be %s, but got %s", "stream 3", string(rec.UserStream))
	}
	if rec.Meta == nil {
		t.Fatalf("Record metadata shouldn't be nil")
	}
}

func BenchmarkContextSetTimer(b *testing.B) {
	ctx := NewContext(&proxy{})
	t1 := time.Now()
	for i := 0; i < b.N; i++ {
		ctx.SetTimer(t1, "test")
	}
}

func BenchmarkContextProduceRecord(b *testing.B) {
	ctx := NewContext(&proxy{})
	for i := 0; i < b.N; i++ {
		ctx.ProduceRecord("stream", "key", "value")
	}
}

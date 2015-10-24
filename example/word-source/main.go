package main

import (
	"errors"
	"github.com/concord/concord-go"
	"log"
	"math/rand"
	"time"
)

const timerName = "main_loop"

type Computation struct {
	words []string
}

// Initialize computation and set timer
func (w *Computation) Init(ctx *concord.Context) error {
	log.Println("Word producer started")
	rand.Seed(time.Now().UnixNano())
	ctx.SetTimer(time.Now(), timerName)
	return nil
}

// Metadata defines name of computation and names of
// input/output stream
func (w *Computation) Metadata() *concord.Metadata {
	return &concord.Metadata{
		Name:    "word-source",
		Outputs: []string{"words"},
	}
}

func (w *Computation) ProcessTimer(ctx *concord.Context, t int64, timerName string) error {
	for i := 0; i < 1000; i++ {
		i := rand.Intn(len(w.words))
		randWord := w.words[i]
		ctx.ProduceRecord("words", randWord, "")
	}
	m := time.Duration(500) * time.Millisecond
	ctx.SetTimer(time.Now().Add(m), timerName)
	return nil
}

func (w *Computation) ProcessRecords(ctx *concord.Context, d interface{}) error {
	return errors.New("ProcessRecords is not implemented")
}

func main() {
	w := &Computation{
		words: []string{"foo", "bar", "baz", "fiz", "buzz"},
	}
	log.Fatal(concord.Serve(w))
}

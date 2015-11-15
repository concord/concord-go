package main

import (
	"errors"
	"github.com/concord/concord-go"
	"log"
	"math/rand"
	"os"
	"time"
)

const timerName = "main_loop"

var logger = log.New(os.Stdout, "", log.LstdFlags)

type Computation struct {
	words []string
}

// Initialize computation and set timer
func (w *Computation) Init(ctx *concord.Context) error {
	logger.Println("Word producer started")
	rand.Seed(time.Now().UnixNano())
	ctx.SetTimer(timerName, time.Now())
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
	for i := 0; i < 100000; i++ {
		i := rand.Intn(len(w.words))
		randWord := w.words[i]
		ctx.ProduceRecord("words", randWord, "")
	}
	logger.Println(time.Now(), "generated random words")
	ctx.SetTimer(timerName, time.Now())
	return nil
}

func (w *Computation) ProcessRecords(ctx *concord.Context, r *concord.Record) error {
	return errors.New("ProcessRecords is not implemented")
}

func main() {
	w := &Computation{
		words: []string{"foo", "bar", "baz", "fiz", "buzz"},
	}
	log.Fatal(concord.Serve(w))
}

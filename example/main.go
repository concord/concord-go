package main

import (
	"github.com/concord/concord-go"
	"log"
	"strings"
	"time"
)

type WordProducer struct {
	words []string
}

func (w *WordProducer) Init(ctx *Context) error {
	log.Println("Word Producer started")
	w.words = []string{"one", "two", "three"}
	ctx.SetTimer(time.Now(), "loop")
	return nil
}

func (w *WordProducer) ProcessRecords(ctx *Context, data interface{}) error {
	return nil
}

func (w *WordProducer) ProcessTimer(ctx *Context, name string) error {
	err := ctx.ProduceRecord(name, "key", "value")
	return err
}

func (w *WordProducer) Metadata() *concord.Metadata {
	return &concord.Metadata{
		Name:    "word-source",
		Inputs:  []string{},
		Outputs: []string{"word"},
	}
}

var comp WordProducer

func main() {
	log.Fatal(concord.Serve(comp))
}

package main

import (
	"errors"
	"github.com/concord/concord-go"
	"log"
)

type Computation struct {
	dict map[string]int
	pidx int
}

func (c *Computation) Init(ctx *concord.Context) error {
	log.Println("Counter initialized")
	return nil
}

func (c *Computation) ProcessTimer(ctx *concord.Context, t int64, timerName string) error {
	return errors.New("Not implemented")
}

func (c *Computation) ProcessRecords(ctx *concord.Context, r *concord.Record) error {
	c.pidx += 1
	k := string(r.Key)
	if _, ok := c.dict[k]; !ok {
		c.dict[k] = 0
	}
	c.dict[k] += 1

	if c.pidx%1024 == 0 {
		log.Println(c.dict)
	}

	return nil
}

func (c *Computation) Metadata() *concord.Metadata {
	return &concord.Metadata{
		Name:   "word-source",
		Inputs: []concord.Stream{concord.NewDefaultStream("words")},
	}
}

func main() {
	comp := &Computation{
		dict: make(map[string]int),
	}
	log.Fatal(concord.Serve(comp))
}

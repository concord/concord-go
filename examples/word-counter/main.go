package main

import (
	"errors"
	"github.com/concord/concord-go"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

type Computation struct {
	dict map[string]int
	pidx int
}

func (c *Computation) Init(ctx *concord.Context) error {
	logger.Println("Counter initialized")
	return nil
}

func (c *Computation) Destroy() error {
	logger.Println("Counter Destroyed")
	return nil
}

func (c *Computation) ProcessTimer(ctx *concord.Context, t int64, timerName string) error {
	return errors.New("Not implemented")
}

func (c *Computation) ProcessRecords(ctx *concord.Context, r *concord.Record) error {
	k := string(r.Key)
	c.pidx += 1
	c.dict[k] += 1

	if c.pidx%100000 == 0 {
		logger.Println(c.dict)
	}

	return nil
}

func (c *Computation) Metadata() *concord.Metadata {
	return &concord.Metadata{
		Name:   "word-counter",
		Inputs: []*concord.Stream{concord.NewStream("words", concord.GroupBy)},
	}
}

func main() {
	comp := &Computation{
		dict: make(map[string]int),
	}
	log.Fatal(concord.Serve(comp))
}

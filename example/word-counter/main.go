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

// class WordCounter(Computation):
//     def __init__(self):
//         self.dict = {}
//         self.pidx = 0 # print index

//     def init(self, ctx):
//         self.concord_logger.info("Counter initialized")

//     def process_timer(self, ctx, key, time):
//         raise Exception('process_timer not implemented')

//     def process_record(self, ctx, record):
//         self.pidx += 1
//         if self.dict.has_key(record.key):
//             self.dict[record.key] += 1
//         else:
//             self.dict[record.key] = 1

//         if (self.pidx % 1024) == 0:
//             self.concord_logger.info(self.dict)

//     def metadata(self):
//         return Metadata(
//             name='word-counter',
//             istreams=['words'],
//             ostreams=[])

// serve_computation(WordCounter())

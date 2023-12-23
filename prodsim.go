package prodsim

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type item int

// Work represents work done at a stage.
type Work func(d, stddev time.Duration)

// ProductionLine represents an imaginary production line.
// Work is done at stages represented by worker functions.
type ProductionLine struct {
	Logger  log.Logger
	Verbose bool
	Stages  []workerFn

	output <-chan item
	ctx    context.Context
}

type Stage struct{}

type workerFn func(context.Context, <-chan item, chan<- item)

func (pl *ProductionLine) AddStage(name string, worker workerFn) {
	pl.Stages = append(pl.Stages, worker)
}

// Run simulates running all stages of the production line.
// It defines an initial and further stages of the line.
func (pl *ProductionLine) Start() {
	prev := make(chan item, 1)

	go func(ctx context.Context, ch chan<- item) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("first stage cancelled!")
				close(ch)
				return
			default:
				ch <- item(i)
				i++
			}
		}
	}(pl.ctx, prev)

	for _, stage := range pl.Stages {
		out := make(chan item, 1)
		go stage(pl.ctx, prev, out)
		prev = out
	}
	pl.output = prev
}

func (pl *ProductionLine) Items() <-chan item {
	return pl.output
}

// newDummyStage takes time and standard deviation and returns
// a worker function. The worker function represents a chunk
// of work that takes time (N). The work can be delayed due to some
// disruptions. Disruptions are represented by deviation.
func newDummyStage(t, stddev time.Duration) workerFn {
	return func(ctx context.Context, in <-chan item, out chan<- item) {
		for item := range in {
			select {
			case <-ctx.Done():
				fmt.Println("worker cancelled!")
				close(out)
				return
			default:
			}

			delay := t + time.Duration(rand.NormFloat64()*float64(stddev))
			time.Sleep(delay)
			out <- item
		}
	}
}

func Run() {
	// todo implement various cancellations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pl := ProductionLine{
		Logger:  *log.Default(),
		Verbose: true,
		ctx:     ctx,
	}

	pl.AddStage("baking", newDummyStage(time.Second, 200*time.Millisecond))
	pl.AddStage("icing", newDummyStage(time.Second, 200*time.Millisecond))
	pl.AddStage("inscribing", newDummyStage(time.Second, 200*time.Millisecond))
	pl.AddStage("packaging", newDummyStage(time.Second, 200*time.Millisecond))

	pl.Start()

	for item := range pl.Items() {
		fmt.Println(item)
	}
}

package service

import (
	"context"
	"fmt"
	"sync"

	"testtask/model"
)

// Worker is struct that implement processing data.
type Worker struct {
	id     int
	ctx    context.Context
	filter *model.Filter
	inCh   chan *model.BasicTitle
	outCh  chan *model.BasicTitle
	wg     *sync.WaitGroup
}

// NewWorker return new instance of Worker.
func NewWorker(id int, ctx context.Context, filter *model.Filter, inCh, outCh chan *model.BasicTitle, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:     id,
		ctx:    ctx,
		filter: filter,
		inCh:   inCh,
		outCh:  outCh,
		wg:     wg,
	}
}

// Start  runs worker queue , wait data from channel filter it and send output to result channel
func (w *Worker) Start() {
	w.wg.Add(1)
	defer w.wg.Done()
	for {
		select {
		case in, ok := <-w.inCh:
			if !ok {
				return
			}
			w.parse(in)
		case <-w.ctx.Done():
			fmt.Printf("stop worker %d\n", w.id)
			return
		}
	}

}

func (w *Worker) parse(in *model.BasicTitle) {
	if in.IsGenres(w.filter.Genre) &&
		in.IsPrimaryTitle(w.filter.PrimaryTitle) {
		w.outCh <- in
	}
}

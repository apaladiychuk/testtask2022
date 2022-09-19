package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/urfave/cli/v2"
	"testtask/model"
	"testtask/service"
)

func Run(ctx *cli.Context, cfg *Config) error {
	inCh := make(chan *model.BasicTitle, 100)
	outCh := make(chan *model.BasicTitle)
	sysSygnal := make(chan os.Signal)
	signal.Notify(sysSygnal, syscall.SIGTERM, syscall.SIGINT)

	workContext, cancel := context.WithCancel(ctx.Context)
	go func() {
		<-sysSygnal
		fmt.Println("Cancel by user....")
		cancel()
	}()
	wg := sync.WaitGroup{}

	for i := 0; i < cfg.amountWorker; i++ {
		worker := service.NewWorker(i, workContext, cfg.Filter, inCh, outCh, &wg)
		go worker.Start()
	}
	r := service.NewReader(cfg.FilePath, inCh)
	go result(workContext, outCh)
	if err := r.Read(workContext); err != nil {
		fmt.Println(err.Error())
	}
	wg.Wait()

	return nil
}

func result(ctx context.Context, outCh chan *model.BasicTitle) {
	fmt.Println("IMDB_ID\t\t|\tTitle\t\t|\t Plot")
	for {
		select {
		case result, ok := <-outCh:
			if !ok {
				return
			}
			fmt.Printf("%s\t|\t%s\t\t|\t%s\n", result.Tconst, result.OriginalTitle, "")
		case <-ctx.Done():
			fmt.Printf("stop writer \n")
			return
		}
	}
}

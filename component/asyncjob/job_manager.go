package asyncjob

import (
	"context"
	"log"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool // cho phép các job chạy tuần tự hay là song song với nhau
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
	}
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))
	errChan := make(chan error, len(g.jobs)) // make a channel to store job's state
	for _, job := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(job)
			continue
		}
		errChan <- g.runJob(ctx, job)
		g.wg.Done()
	}
	g.wg.Wait()
	var err error
	for range g.jobs {
		if v := <-errChan; v != nil {
			err = v
		}
	}
	log.Println("Done Group")
	return err
}

func (g *group) runJob(ctx context.Context, j Job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			log.Println(err)
			if j.Retry(ctx) == nil {
				return nil
			}
			// if j.State() == StateRetryFailed {
			// 	log.Println("--->>> retry failed")
			// 	return err
			// }
		}
	}
	return nil
}

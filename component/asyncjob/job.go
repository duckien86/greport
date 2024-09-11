package asyncjob

import (
	"context"
	"errors"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(time []time.Duration)
}

const (
	defaultMaXTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 3}
)

type JobHandler func(ctx context.Context) error
type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeOut
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaXTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		stopChan:   make(chan bool),
	}
	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning
	err := j.handler(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}
	j.state = StateCompleted
	return nil
}
func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	if j.retryIndex == len(j.config.Retries) {
		j.state = StateFailed
		return errors.New("reach retry times")
	}
	time.Sleep(j.config.Retries[j.retryIndex])
	err := j.Execute(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}
func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }
func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}
	j.config.Retries = times
}

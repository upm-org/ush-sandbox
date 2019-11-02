package interp

import (
	"context"
	"errors"
	"sync"
)

type asyncable interface {
	Pause() error
	Unpause() error
	Wait()
}

var (
	pausedErr   = errors.New("can't pause paused context")
	unPausedErr = errors.New("can't unpause not paused context")
)

type asyncCtx struct {
	context.Context

	mu        sync.Mutex
	done      chan struct{}
	err       error
	pauseChan chan struct{}
	isPaused  bool
}

func (a *asyncCtx) toggle() {
	go func() {
		a.pauseChan <- struct{}{}
	}()
	a.isPaused = !a.isPaused
}

func (a *asyncCtx) Pause() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isPaused {
		return pausedErr
	}

	a.toggle()
	return nil
}

func (a *asyncCtx) Unpause() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isPaused {
		return unPausedErr
	}

	a.toggle()
	return nil
}

func (a *asyncCtx) Wait() {
	select {
	case <-a.pauseChan:
		<-a.pauseChan
		return
	default:
		return
	}
}

func newAsyncCtx(parent context.Context) asyncable {
	return &asyncCtx{Context: parent, pauseChan: make(chan struct{})}
}

func WithAsync(parent context.Context) (context.Context, context.CancelFunc) {
	c, cancelFunc := context.WithCancel(parent)
	return newAsyncCtx(c).(context.Context), cancelFunc
}

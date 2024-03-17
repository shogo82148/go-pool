package pool

import (
	"context"
	"sync"
)

type Pool[T any] struct {
	New func(ctx context.Context) (T, error)

	mu        sync.Mutex
	freeItems []T
}

func (p *Pool[T]) Get(ctx context.Context) (T, error) {
	p.mu.Lock()
	if len(p.freeItems) > 0 {
		l := len(p.freeItems)
		t := p.freeItems[l-1]
		p.freeItems = p.freeItems[:l-1]
		p.mu.Unlock()
		return t, nil
	}
	p.mu.Unlock()
	return p.New(ctx)
}

func (p *Pool[T]) Put(x T) {
	p.mu.Lock()
	p.freeItems = append(p.freeItems, x)
	p.mu.Unlock()
}

func (p *Pool[T]) Close() error {
	return nil
}

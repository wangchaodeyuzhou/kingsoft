package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type pool struct {
	m        sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

var ErrClosePool = errors.New("Pool has been closed. ")

func New(fn func() (io.Closer, error), size uint) (*pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
	}, nil
}

func (p *pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		if !ok {
			return nil, ErrClosePool
		}
		return r, nil
	default:
		fmt.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

func (p *pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resource <- r:
		fmt.Println("Release:", "in queue")
	default:
		fmt.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resource)
	for r := range p.resource {
		r.Close()
	}
}

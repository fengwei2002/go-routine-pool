package pool

import "sync"

type Pool struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewPool(maxSize int) *Pool {
	return &Pool{
		c:  make(chan struct{}, maxSize),
		wg: new(sync.WaitGroup),
	}
}

func (this *Pool) Add(delta int) {
	this.wg.Add(delta)
	for i := 0; i < delta; i++ {
		this.c <- struct{}{}
	}
}

func (this *Pool) Done() {
	<-this.c
	this.wg.Done()
}

func (this *Pool) Wait() {
	this.wg.Wait()
}

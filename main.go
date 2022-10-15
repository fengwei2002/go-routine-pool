package main

import (
	"fmt"
	"pool"
	"time"
)

var p = pool.NewPool(10)

func main() {
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Work(i)
	}
	p.Wait()
}

func Work(i int) {
	defer p.Done()
	p.Add(1)
	time.Sleep(time.Second)
	fmt.Println(i)
}

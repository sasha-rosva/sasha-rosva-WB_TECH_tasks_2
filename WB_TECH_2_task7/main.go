package main

import (
	"fmt"
	"sync"
	"time"
)

func manyToOne(channels ...<-chan interface{}) <-chan interface{} {
	var group sync.WaitGroup
	output := make(chan interface{}, 1)
	for i := range channels {
		group.Add(1)
		go func(input <-chan interface{}) {
			for val := range input {
				output <- val
			}
			group.Done()
		}(channels[i])
	}
	go func() {
		group.Wait()
		close(output)
	}()
	return output
}
func main() {
	sig := func(second int, n int) <-chan interface{} {
		c := make(chan interface{}, 1)
		go func() {
			defer close(c)
			for i := 0; i < second; i++ {
				c <- n
				time.Sleep(time.Second)
			}

		}()
		return c
	}
	start := time.Now()
	out := manyToOne(
		sig(20, 0),
		sig(15, 1),
		sig(10, 2),
		sig(5, 3),
		sig(1, 4),
	)
	for vvv := range out {
		fmt.Println(vvv)
	}
	fmt.Printf("fone after %v", time.Since(start))
}
